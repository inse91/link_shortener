package service

import (
	"errors"
	"hash/crc64"
	errs "link_shortener/internal/error"
	"link_shortener/internal/model"
	store2 "link_shortener/internal/store"
	"log"
	"strings"
)

type Shorter struct {
	log   *log.Logger
	store store2.Store
}

func NewShorter(store store2.Store, log *log.Logger) *Shorter {
	return &Shorter{
		log:   log,
		store: store,
	}
}

func (s Shorter) Get(short string) (full string, err error) {

	full, err = s.store.Get(short)
	if err != nil {
		s.log.Printf("failed getting full from %s\n", short)
		return "", err
	}

	return full, nil
}

func (s Shorter) Create(full string) (short string, err error) {

	hash := getHash(full)
	short = base63Encode(hash, model.AlphabetBase63)

	// проверка на коллизии
	var isOk bool
	//for isOk, err = s.collisionCheck(full, short); ; {
	for isOk, err = s.collisionCheck(full, short); !errors.Is(err, errs.ErrNotFound); {

		if err == nil && isOk {
			return short, nil
		}

		//if errors.Is(err, errs.ErrNotFound) {
		//	break
		//}

		short += string(model.AlphabetBase63[0])
	}

	err = s.store.Create(short, full)
	if err != nil {
		return "", err
	}

	return short, nil
}

func (s Shorter) collisionCheck(full, short string) (bool, error) {
	existingFull, err := s.Get(short)
	if err != nil {
		return false, err
	}
	return existingFull == full, nil
}

func getHash(src string) (hash uint64) {
	return crc64.Checksum([]byte(src), model.HashTable)
}

func base63Encode(n uint64, alphabet string) string {
	base := uint64(len(alphabet))
	if n == 0 {
		return string(alphabet[0])
	}
	chars := []string{}
	for n > 0 {
		r := n % base
		chars = append(chars, string(alphabet[r]))
		n = n / base
	}
	// Reverse the order
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return strings.Join(chars, "")
}
