package models

type Answer struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Question Question `json:"question"`
	QuestionID int `json:"question_id"`
	Correct bool `json:"correct"`
}
