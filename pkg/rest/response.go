package rest

type SuccessfulResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
