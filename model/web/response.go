package web

type Response struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	RecordsTotal int         `json:"recordsTotal"`
	Data         interface{} `json:"data"`
}
