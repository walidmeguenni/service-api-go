package types

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any      `json:"data"`
}
