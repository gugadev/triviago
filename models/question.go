package models

type Question struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Category Category
	CategoryID int `json:"category_id"`
}
