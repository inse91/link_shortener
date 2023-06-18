package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "link_shortener/api"
	errs "link_shortener/internal/error"
	"link_shortener/internal/model"
	"link_shortener/internal/service"
	"log"
	"net/http"
)

type ShorterHandler struct {
	logger  *log.Logger
	service *service.Shorter
}

func NewHttp(srv *service.Shorter, log *log.Logger) *ShorterHandler {
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

	r.Post("/Create", sh.Create)
	r.Post("/Get", sh.Get)
}

// Get godoc
// @Summary      возвращает исходную ссылку
// @Description  возвращает исходную ссылку
// @Accept       json
// @Produce      json
// @Param        link   body      string  true  "Сокрещенная ссылка"
// @Success      200  {object}   model.LinkResponse
// @Failure      404  {object}   model.LinkResponse
// @Failure      500  {object}   model.LinkResponse
// @Router       /get [post]
func (sh *ShorterHandler) Get(c *fiber.Ctx) (err error) {
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
		if err == errs.ErrNotFound {
			if err = c.
				Status(http.StatusNotFound).
				JSON(model.LinkResponse{
					Success: false,
					Error:   err.Error(),
				}); err != nil {
				sh.logger.Printf("failed sending response: %s", err.Error())
			}
			return
		}
		if err = c.
			Status(http.StatusInternalServerError).
			JSON(model.LinkResponse{
				Success: false,
				Error:   err.Error(),
			}); err != nil {
			sh.logger.Printf("failed sending response: %s", err.Error())
		}
		return
	}

	if err = c.
		Status(http.StatusOK).
		JSON(model.LinkResponse{
			Link:    model.Link{Link: fullLink},
			Success: true,
		}); err != nil {
		sh.logger.Printf("failed sending response: %s", err.Error())
	}

	return
}

// Create godoc
// @Summary      создает сокращенную весию ссылки
// @Description  создает сокращенную весию ссылки
// @Accept       json
// @Produce      json
// @Param        link   body      string  true  "Исходная ссылка"
// @Success      200  {object}   model.LinkResponse
// @Failure      500  {object}   model.LinkResponse
// @Router       /create [post]
func (sh *ShorterHandler) Create(c *fiber.Ctx) (err error) {
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
		return c.JSON(model.LinkResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(model.LinkResponse{
		Link:    model.Link{Link: newShortLink},
		Success: true,
	})
}
