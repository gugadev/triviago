package responses


type Stat struct {
	Category string `json:"category"`
	Total int `json:"total"`
	Hits int `json:"hits"`
}
