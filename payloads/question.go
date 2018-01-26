package payloads

import "github.com/gugadev/triviago/models"

// NewQuestion payload of a POST create question
type NewQuestion struct {
	Question models.Question `json:"question"`
	Answers  []models.Answer `json:"answers"`
}
