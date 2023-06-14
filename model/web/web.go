package web

const InternalServerErrorMessage = "Unexpected error occurred"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
