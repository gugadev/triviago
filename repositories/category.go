package repositories

import (
	"github.com/gugadev/triviago/database"
	"github.com/gugadev/triviago/models"
)

type CategoryRepo struct{}

func (r *CategoryRepo) All() []models.Category {
	var resultset []models.Category

	db := database.Connect()
	defer db.Close()

	db.Find(&resultset)
	return resultset
}
