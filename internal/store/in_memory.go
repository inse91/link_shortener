package store

import (
	errs "link_shortener/internal/error"
	"log"
)

type inMem struct {
	store map[string]string
	log   *log.Logger
}

func (i *inMem) Create(sh string, long string) error {
	i.store[sh] = long
	return nil
}

func (i *inMem) Get(short string) (string, error) {
	if v, ok := i.store[short]; ok {
		return v, nil
	}
	return "", errs.ErrNotFound
}

func newInMem(log *log.Logger) Store {
	return &inMem{
		log:   log,
		store: make(map[string]string),
	}
}
