package main

import (
	"github.com/dyfun/go-echo-fw/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/healt-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostHandler)
	e.GET("/post/:id", handlers.SingleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
