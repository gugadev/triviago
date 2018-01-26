package repositories

import (
	"github.com/gugadev/triviago/database"
	"github.com/gugadev/triviago/models"
)

// CategoryRepo repo for categories
type CategoryRepo struct{}

/*
All returns all categories
*/
func (r *CategoryRepo) All() []models.Category {
	var resultset []models.Category

	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	db.Find(&resultset)
	return resultset
}
