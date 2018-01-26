package handlers

import (
	"github.com/labstack/echo"
	"github.com/gugadev/triviago/repositories"
)

func GetCategories(c echo.Context) error {
	repo := repositories.CategoryRepo{}
	return c.JSON(200, repo.All())
}
