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

func newPg(ctx context.Context, dbConn string, log *log.Logger) (Store, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		"postgres", "password", dbConn, "postgres",
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
	p.log.Printf("creating record %s => %s\n", short, full)
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

	return nil
}

func (p *pg) Get(short string) (string, error) {
	p.log.Printf("getting full link form %s\n", short)
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
			return "", errs.ErrNotFound
		}
		return "", err
	}

	return full, nil
}
