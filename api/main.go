package main

import (
	"net/http"

	"github.com/ThyLeader/neural-api/generation"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var nets = map[string]string{
	"shakespeare": "../generation/shakespeare.lstm",
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", root)
	e.GET("/generate/:type", gen)

	e.Logger.Fatal(e.Start(":1323"))
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func gen(c echo.Context) error {
	t := c.Param("type")
	l := c.QueryParam("length")

	if l == "" || t == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	f, ok := nets[t]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Type Not Found")
	}

	g, err := generation.Generate(f, l)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"data": g,
		},
	)
}
