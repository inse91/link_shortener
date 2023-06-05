package handler

import (
	"github.com/gofiber/fiber/v2"
	"link_shortener/internal/model"
	"link_shortener/internal/service"
	"log"
	"net/http"
)

type ShorterHandler struct {
	logger  *log.Logger
	service *service.Shorter
}

func NewShorterHandler(srv *service.Shorter, log *log.Logger) *ShorterHandler {
	return &ShorterHandler{
		logger:  log,
		service: srv,
	}
}

func (sh *ShorterHandler) Register(r *fiber.App) {
	r.Get("/hello", func(c *fiber.Ctx) error {

		err := c.SendString("hello world")
		if err != nil {
			return err
		}
		return nil
	})

	r.Post("/create", func(c *fiber.Ctx) (err error) {

		c.Accepts("application/json")

		link := new(model.Link)
		if err = c.BodyParser(link); err != nil {
			if err = c.
				Status(http.StatusBadRequest).
				JSON(model.LinkResponse{
					Success: false,
					Error:   err.Error(),
				}); err != nil {
				sh.logger.Println(err)
			}
			return
		}

		var newShortLink string
		if newShortLink, err = sh.service.Create(link.Link); err != nil {
			if err = c.JSON(model.LinkResponse{
				Success: false,
				Error:   err.Error(),
			}); err != nil {
				sh.logger.Println(err)
			}
			return
		}

		if err = c.JSON(model.LinkResponse{
			Link:    model.Link{Link: newShortLink},
			Success: true,
		}); err != nil {
			return
		}

		return
	})

	r.Post("/get", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")

		link := new(model.Link)
		if err = c.BodyParser(link); err != nil {
			if err = c.
				Status(http.StatusBadRequest).
				JSON(model.LinkResponse{
					Success: false,
					Error:   err.Error(),
				}); err != nil {
				sh.logger.Println(err)
			}
			return
		}

		var fullLink string
		if fullLink, err = sh.service.Get(link.Link); err != nil {
			if err = c.
				Status(http.StatusInternalServerError).
				JSON(model.LinkResponse{
					Success: false,
					Error:   err.Error(),
				}); err != nil {
				sh.logger.Println(err)
			}
			return
		}

		if err = c.
			Status(http.StatusOK).
			JSON(model.LinkResponse{
				Link:    model.Link{Link: fullLink},
				Success: true,
			}); err != nil {
			sh.logger.Println(err)
		}

		return
	})
}
