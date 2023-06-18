package handler

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	errs "link_shortener/internal/error"
	"link_shortener/internal/service"
	"log"
	"time"
)

type Shortener struct {
	logger  *log.Logger
	service *service.Shorter
}

func (s Shortener) Create(ctx context.Context, request *Request) (*Response, error) {
	link := request.GetLink()
	shortedLink, err := s.service.Create(link)
	if err != nil {
		return &Response{
			Success: false,
			Link:    "",
		}, nil
	}
	return &Response{
		Link:    shortedLink,
		Success: true,
	}, nil
}

func (s Shortener) Get(ctx context.Context, request *Request) (*Response, error) {
	link := request.GetLink()

	fullLink, err := s.service.Get(link)
	if err != nil {
		if err == errs.ErrNotFound {
			return &Response{
				Success: false,
				Link:    err.Error(),
			}, nil
		}
		return &Response{
			Success: false,
			Link:    "internal error: " + err.Error(),
		}, nil

	}
	return &Response{
		Link:    fullLink,
		Success: true,
	}, nil
}

func (s Shortener) mustEmbedUnimplementedShortenerServer() {
}

func NewGrpc(logger *log.Logger, service *service.Shorter) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	RegisterShortenerServer(grpcServer, &Shortener{
		logger:  logger,
		service: service,
	})

	return grpcServer
}
