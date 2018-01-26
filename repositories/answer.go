package repositories

import (
	"github.com/gugadev/triviago/database" // nolint: gotype
	"github.com/gugadev/triviago/models"
)

// AnswerRepo repo for Answer model
type AnswerRepo struct{}

/*
Create create a new anser
*/
func (r *AnswerRepo) Create(answer models.Answer) {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	db.Create(&answer)
}

/*
GetFor get all answers for a question
*/
func (r *AnswerRepo) GetFor(questionID int) []models.Answer {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	var answers []models.Answer

	db.Where("question_id = ?", questionID).Find(&answers)
	return answers
}

/*
Find get an answer by it's ID
*/
func (r *AnswerRepo) Find(ID int) models.Answer {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	var answer models.Answer
	db.First(&answer, ID)
	return answer
}
