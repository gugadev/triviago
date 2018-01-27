package handlers

import (
	"github.com/gugadev/triviago/repositories"
	"github.com/labstack/echo"
)

/*
GetCategories get all categories
*/
func GetCategories(c echo.Context) error {
	repo := repositories.CategoryRepo{}
	return c.JSON(200, repo.All())
}
