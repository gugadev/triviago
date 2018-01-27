package responses

// Stat http response for stats
type Stat struct {
	Category string `json:"category"`
	Total    int    `json:"total"`
	Hits     int    `json:"hits"`
}
