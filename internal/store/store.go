package store

import (
	"context"
	"log"
)

type Store interface {
	Create(string, string) error
	Get(string) (string, error)
	//Ping() error
}

func New(ctx context.Context, dbConn string, logger *log.Logger) (Store, error) {
	if dbConn == "" {
		return newInMem(logger), nil
	}
	return newPg(ctx, dbConn, logger)
}
