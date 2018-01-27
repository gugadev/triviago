package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gugadev/triviago/handlers" // nolint: gotype
	// nolint: gotype
	"github.com/gobuffalo/packr"
	"github.com/labstack/echo"
)

func main() {
	// setup.Migrate()
	// setup.Populate()

	app := echo.New()

	assetHandler := http.FileServer(packr.NewBox("./public"))

	app.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static", assetHandler)))

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
