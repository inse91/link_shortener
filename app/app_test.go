package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"link_shortener/internal/handler"
	"link_shortener/internal/model"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	httpPort = "8081"
	grpcPort = "8082"
	dbConn   = ""
	//dbConn   = "localhost:5432"
)

func TestE2E_AUTO(t *testing.T) {
	setup(t)
	go Start()

	testLink := "google.com/test/api"
	fatal := t.Fatal

	time.Sleep(100 * time.Millisecond)

	t.Run("create via http + get via grpc", func(t *testing.T) {

		bts, err := json.Marshal(model.Link{Link: testLink})
		if err != nil {
			fatal(err)
		}
		response, err := http.DefaultClient.Post(
			fmt.Sprintf("http://localhost:%s/create", httpPort),
			"application/json",
			bytes.NewReader(bts),
		)
		if err != nil {
			t.Fatal(err)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			t.Fatal("resp code is" + strconv.Itoa(response.StatusCode))
		}

		respData := new(model.LinkResponse)
		if err = json.NewDecoder(response.Body).Decode(respData); err != nil {
			t.Fatal(err)
		}

		if !respData.Success {
			t.Fatal("resp is not success: " + respData.Error)
		}

		log.Printf("got short link: %s", respData.Link.Link)

		// grpc
		conn, err := grpc.Dial(":"+grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Fatal(err)
		}

		gResp, err := handler.NewShortenerClient(conn).Get(context.Background(), &handler.Request{
			Link: respData.Link.Link,
		})
		if err != nil {
			t.Fatal(err)
		}

		if !gResp.GetSuccess() {
			t.Fatal("grpc not success resp")
		}

		log.Printf("got full link: %s", gResp.GetLink())

		if gResp.GetLink() != testLink {
			t.Fatal("link are not equal")
		}

	})

	t.Run("create via grpc + get via http", func(t *testing.T) {

		// grpc
		conn, err := grpc.Dial(":"+grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Fatal(err)
		}

		gResp, err := handler.NewShortenerClient(conn).Create(context.Background(), &handler.Request{
			Link: testLink,
		})
		if err != nil {
			t.Fatal(err)
		}

		if !gResp.GetSuccess() {
			t.Fatal("grpc not success resp")
		}

		log.Printf("got short link: %s", gResp.GetLink())

		bts, err := json.Marshal(model.Link{Link: gResp.GetLink()})
		if err != nil {
			fatal(err)
		}
		response, err := http.DefaultClient.Post(
			fmt.Sprintf("http://localhost:%s/get", httpPort),
			"application/json",
			bytes.NewReader(bts),
		)
		if err != nil {
			t.Fatal(err)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			t.Fatal("resp code is " + strconv.Itoa(response.StatusCode))
		}

		respData := new(model.LinkResponse)
		if err = json.NewDecoder(response.Body).Decode(respData); err != nil {
			t.Fatal(err)
		}

		if !respData.Success {
			t.Fatal("resp is not success: " + respData.Error)
		}

		log.Printf("got full link: %s", respData.Link.Link)

	})

	t.Run("unsuccessful http", func(t *testing.T) {
		bts, err := json.Marshal(model.Link{Link: testLink + "FOO"})
		if err != nil {
			fatal(err)
		}
		response, err := http.DefaultClient.Post(
			fmt.Sprintf("http://localhost:%s/get", httpPort),
			"application/json",
			bytes.NewReader(bts),
		)
		if err != nil {
			fatal(err)
		}

		defer response.Body.Close()

		respData := new(model.LinkResponse)
		if err = json.NewDecoder(response.Body).Decode(respData); err != nil {
			t.Fatal(err)
		}

		if respData.Success || response.StatusCode != http.StatusNotFound {
			t.Fatal("resp is success")
		}
	})

	t.Run("unsuccessful grpc", func(t *testing.T) {

		// grpc
		conn, err := grpc.Dial(":"+grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Fatal(err)
		}

		gResp, err := handler.NewShortenerClient(conn).Get(context.Background(), &handler.Request{
			Link: testLink + "FOO",
		})
		if err != nil {
			t.Fatal(err)
		}

		if gResp.GetSuccess() {
			t.Fatal("grpc resp is success")
		}

	})

}

func TestApp(t *testing.T) {
	setup(t)
	Start()
}

func setup(t *testing.T) {

	err := os.Setenv("HTTP_PORT", httpPort)
	if err != nil {
		t.Fatal(err)
	}

	if err = os.Setenv("GRPC_PORT", grpcPort); err != nil {
		t.Fatal(err)
	}

	if err = os.Setenv("DB_CONNECTION", dbConn); err != nil {
		t.Fatal(err)
	}

}
