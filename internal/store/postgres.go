package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	errs "link_shortener/internal/error"
	"log"
)

type pg struct {
	db     *sql.DB
	client *pgx.Conn
	log    *log.Logger
}

func newPg(ctx context.Context, connPort string, log *log.Logger) (Store, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"postgres", "password", "localhost", connPort, "postgres",
	)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &pg{
		db:     nil,
		log:    log,
		client: conn,
	}, nil
}

func (p *pg) Create(short, full string) error {

	q := `INSERT
			INTO public.links (short, "full") 
			VALUES ($1, $2)
			RETURNING id`
	var id string
	if err := p.client.
		QueryRow(context.TODO(), q, short, full).
		Scan(&id); err != nil {
		return err
	}

	p.log.Printf("link created with id %s", id)
	return nil
}

func (p *pg) Get(short string) (string, error) {

	q := `
		SELECT ("full")
			FROM public.links 
			WHERE short = $1
		`

	var full string
	if err := p.client.
		QueryRow(context.TODO(), q, short).
		Scan(&full); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.log.Printf("link with short %s not found", short)
			return "", errs.ErrNotFound
		}
		return "", err
	}

	p.log.Printf("found full link %s (from short) %s", full, short)
	return full, nil
}
