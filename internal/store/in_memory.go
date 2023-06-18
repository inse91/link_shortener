package store

import (
	errs "link_shortener/internal/error"
	"log"
	"sync"
)

type inMem struct {
	sync.RWMutex
	store map[string]string
	log   *log.Logger
}

func (im *inMem) Create(short string, full string) error {
	im.log.Printf("creating record %s => %s\n", short, full)
	im.Lock()
	defer im.Unlock()
	im.store[short] = full
	return nil
}

func (im *inMem) Get(short string) (string, error) {
	im.log.Printf("getting full link form %s\n", short)
	im.RLock()
	defer im.RUnlock()
	if full, ok := im.store[short]; ok {
		return full, nil
	}
	return "", errs.ErrNotFound
}

func newInMem(log *log.Logger) Store {
	return &inMem{
		log:   log,
		store: make(map[string]string),
	}
}
