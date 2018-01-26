package responses

import "github.com/gugadev/triviago/models"

type QuestionResponse struct {
	Question models.Question `json:"question"`
	Answers []models.Answer `json:"answers"`
}
