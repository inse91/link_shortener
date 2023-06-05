package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"link_shortener/config"
	"link_shortener/internal/handler"
	"link_shortener/internal/proto"
	"link_shortener/internal/service"
	"link_shortener/internal/store"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start() {

	cfg := config.Get()
	logger := log.Default()

	ctx := context.Background()

	s, err := store.New(ctx, cfg.DataBase, logger)
	if err != nil {
		logger.Fatal(err)
	}

	shorter := service.NewShorter(s, logger)
	h := handler.NewShorterHandler(shorter, logger)

	fiberApp := fiber.New()
	h.Register(fiberApp)

	go func() {
		log.Fatal(fiberApp.Listen(":" + cfg.HttpPort))
	}()

	grpcServer := proto.NewShorterServer(logger, shorter)
	//lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("grpc started lis on port: " + cfg.GrpcPort)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	oscall := <-interrupt
	log.Printf("shutdown server, %s", oscall)
	grpcServer.GracefulStop()
	if err = fiberApp.Shutdown(); err != nil {
		log.Printf("error occured on server shutting down: %v", err)
	}
}
