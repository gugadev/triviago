package models


type Score struct {
	ID int `json:"id"`
	Question Question `json:"question"`
	QuestionID int `json:"question_id"`
	Hit bool `json:"hit"`
}
