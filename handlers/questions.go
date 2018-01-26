package handlers

import (
	"strconv"
	"strings"

	"github.com/gugadev/triviago/helpers" // nolint: gotype
	"github.com/gugadev/triviago/models"
	"github.com/gugadev/triviago/payloads"
	"github.com/gugadev/triviago/repositories"
	"github.com/gugadev/triviago/responses"
	"github.com/labstack/echo"
)

/*
GiveQuestion returns a question
*/
func GiveQuestion(c echo.Context) error {
	repo := new(repositories.QuestionRepo)
	arepo := new(repositories.AnswerRepo)

	query := strings.Split(c.Request().URL.Query().Get("cats"), ",")
	cats := helpers.MapAtoi(query, helpers.ApplyMapAtoi)

	question := repo.Get(cats)
	answers := arepo.GetFor(question.ID)

	return c.JSON(200, responses.QuestionResponse{
		Question: question,
		Answers:  answers,
	})
}

/*
CheckAnswer check if an answer is correct */
func CheckAnswer(c echo.Context) error {
	repo := new(repositories.QuestionRepo)
	arepo := new(repositories.AnswerRepo)
	srepo := new(repositories.ScoreRepo)

	questionID, _ := strconv.Atoi(c.QueryParam("question")) // nolint: gas,errcheck
	answerID, _ := strconv.Atoi(c.QueryParam("answer"))     // nolint: gas,errcheck

	question := repo.Find(questionID)
	answer := arepo.Find(answerID)

	if answer.ID == 0 || question.ID == 0 {
		return c.JSON(400, responses.Answer{
			Message: "Answer or question doesn't exists",
		})
	}
	if answer.QuestionID != question.ID {
		return c.JSON(400, responses.Answer{
			Message: "Answer does not belongs to the question",
		})
	}
	srepo.Insert(models.Score{
		Question: question,
		Hit:      answer.Correct,
	})
	return c.JSON(200, responses.Answer{
		Success: answer.Correct,
	})
}

/*
CreateQuestion creates a new question
*/
func CreateQuestion(c echo.Context) error {
	var payload payloads.NewQuestion
	var questionRepo repositories.QuestionRepo
	var answerRepo repositories.AnswerRepo

	c.Bind(&payload) // nolint: gas,errcheck

	// create the question first
	questionID := questionRepo.Create(payload.Question)

	// create the answers then
	for _, answer := range payload.Answers {
		answer.QuestionID = questionID
		answerRepo.Create(answer)
	}
	return c.String(200, "success")
}
