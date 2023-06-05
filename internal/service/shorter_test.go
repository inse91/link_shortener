package service

import (
	"context"
	"link_shortener/internal/store"
	"log"
	"testing"
)

var sh *Shorter

func init() {
	s, _ := store.New(context.Background(), "", log.Default())
	sh = &Shorter{
		log:   log.Default(),
		store: s,
	}
}

func TestShorter_Create(t *testing.T) {

	links := []string{
		"ya.ru",
		"googole.com",
		"vk.ru/app/test",
		"vk.ru/app/test1",
	}

	newLinks := make([]string, 0, len(links))

	for _, link := range links {
		short, err := sh.Create(link)
		if err != nil {
			t.Fatal(err)
		}
		get, err := sh.Get(short)
		if err != nil {
			t.Fatal(err)
		}
		newLinks = append(newLinks, get)
	}

	for i := range newLinks {
		if newLinks[i] != links[i] {
			t.Fatal("links are not equal")
		}
	}

}
