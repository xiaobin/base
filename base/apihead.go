package base

type RestResponseHead struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
