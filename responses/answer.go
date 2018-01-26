package responses


type Answer struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Success bool `json:"success"`
}
