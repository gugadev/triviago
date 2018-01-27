package handlers

import (
	"github.com/gugadev/triviago/repositories"
	"github.com/labstack/echo"
)

/*
GetStats get stats of the player
*/
func GetStats(c echo.Context) error {
	scoreRepo := new(repositories.ScoreRepo)
	stats := scoreRepo.GetStats()
	return c.JSON(200, stats)
}
