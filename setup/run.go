package setup

import (
	"github.com/gugadev/triviago/database"
	"github.com/gugadev/triviago/models"
)

/*
Migrate recreate the database
*/
func Migrate() {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	db.DropTableIfExists(&models.Category{}, &models.Question{}, &models.Answer{}, models.Score{})
	db.AutoMigrate(&models.Category{}, &models.Question{}, &models.Answer{}, models.Score{})
}

/*
Populate initially inserts
*/
func Populate() {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	db.Create(&models.Category{
		Name: "General",
	})
	db.Create(&models.Category{
		Name: "Ciencia",
	})
	db.Create(&models.Category{
		Name: "Mitología",
	})
	db.Create(&models.Category{
		Name: "Deportes",
	})
	db.Create(&models.Category{
		Name: "Geografía",
	})
	db.Create(&models.Category{
		Name: "Historia",
	})
	db.Create(&models.Category{
		Name: "Arte",
	})
	db.Create(&models.Category{
		Name: "Animales",
	})
}
