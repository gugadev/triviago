package repositories

import (
	"github.com/gugadev/triviago/database"
	"github.com/gugadev/triviago/models"
)

// QuestionRepo repository for Question
type QuestionRepo struct{}

/*
Create create a new question
@returns models.Question recently created
*/
func (r *QuestionRepo) Create(question models.Question) int {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	db.Create(&question)

	return question.ID
}

/*
Get get a question randomly from entered categories
*/
func (r *QuestionRepo) Get(categories []int) models.Question {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	var question models.Question

	db.Where("category_id IN(?)", categories).Order("random()").Limit(1).Find(&question)
	return question
}

/*
Find find a question by ID
*/
func (r *QuestionRepo) Find(ID int) models.Question {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	var question models.Question

	db.First(&question, ID)
	return question
}
