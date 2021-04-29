package base

type SimplePageData struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
