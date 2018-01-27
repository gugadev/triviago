package responses

import "github.com/gugadev/triviago/models"

// QuestionResponse http response for question req
type QuestionResponse struct {
	Question models.Question `json:"question"`
	Answers  []models.Answer `json:"answers"`
}
