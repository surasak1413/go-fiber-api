package response

type Field struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
