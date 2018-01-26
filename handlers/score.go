package handlers

import (
	"github.com/labstack/echo"
	"github.com/gugadev/triviago/repositories"
)

func GetStats(c echo.Context) error {
	scoreRepo := new(repositories.ScoreRepo)
	stats := scoreRepo.GetStats()
	return c.JSON(200, stats)
}
