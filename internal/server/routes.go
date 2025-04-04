package server

import (
	"net/http"
	"wordle-backend/internal/words"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	wg_en *words.WordGen
	wg_sr *words.WordGen
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	wg_en = words.NewWordGen("english")
	wg_sr = words.NewWordGen("serbian")

	e.GET("/en", s.EnglishHandler)
	e.GET("/sr", s.SerbianHandler)

	return e
}

func (s *Server) EnglishHandler(c echo.Context) error {
	resp := wg_en.RandomWord().Word

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) SerbianHandler(c echo.Context) error {
	resp := wg_sr.RandomWord().Word

	return c.JSON(http.StatusOK, resp)
}
