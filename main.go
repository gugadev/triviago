package main

import (
	"io/ioutil"

	"github.com/gugadev/triviago/handlers" // nolint: gotype
	"github.com/gugadev/triviago/setup"    // nolint: gotype
	"github.com/labstack/echo"
)

func main() {
	setup.Migrate()
	setup.Populate()

	app := echo.New()

	app.Static("/static", "client/public")

	app.GET("/api/categories", handlers.GetCategories)
	app.GET("/api/questions", handlers.GiveQuestion)
	app.POST("/api/questions", handlers.CreateQuestion)
	app.POST("/api/answers/check", handlers.CheckAnswer)
	app.GET("/api/score", handlers.GetStats)
	app.GET("/*", func(ctx echo.Context) error {
		c, _ := ioutil.ReadFile("client/index.html") // nolint: gas,errcheck
		return ctx.HTML(200, string(c))
	})

	app.Start("0.0.0.0:6789") // nolint: gas,errcheck
}
