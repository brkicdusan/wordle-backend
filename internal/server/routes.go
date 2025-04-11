package server

import (
	"net/http"
	"strconv"
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
	e.GET("/en/:count", s.EnglishMultiHandler)
	e.GET("/sr/:count", s.SerbianMultiHandler)

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

func (s *Server) EnglishMultiHandler(c echo.Context) error {
	return MultiHandler(wg_en, c)
}

func (s *Server) SerbianMultiHandler(c echo.Context) error {
	return MultiHandler(wg_sr, c)
}

func MultiHandler(wg *words.WordGen, c echo.Context) error {
	countStr := c.Param("count")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid count.")
	}

	list, err := wg.GetN(count)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid count.")
	}

	return c.JSON(http.StatusOK, list)
}
