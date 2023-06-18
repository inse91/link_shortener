package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "link_shortener/api"
	"link_shortener/config"
	"link_shortener/internal/handler"
	"link_shortener/internal/service"
	"link_shortener/internal/store"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// @title Link Shortener API
// @version 0.9
// @description Link Shortener API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email ozhegov@elma-bs.ru
// @host localhost:8080
// @BasePath /

func Start() {

	cfg := config.Get()
	logger := log.Default()
	ctx, cancel := context.WithCancel(context.Background())

	st, err := store.New(ctx, cfg.DataBase, logger)
	if err != nil {
		logger.Fatal(err)
	}

	shorter := service.NewShorter(st, logger)

	// http
	h := handler.NewHttp(shorter, logger)
	fiberApp := fiber.New()
	h.Register(fiberApp)
	fiberApp.Get("/swagger/*", swagger.HandlerDefault)

	go func() {
		log.Fatal(fiberApp.Listen(":" + cfg.HttpPort))
	}()

	// grpc
	grpcServer := handler.NewGrpc(logger, shorter)
	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		logger.Fatal(err)
	}
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	oscall := <-interrupt
	log.Printf("shutdown server, %st", oscall)
	cancel()
	grpcServer.GracefulStop()
	fiberApp.Shutdown()
}
