package base

type SimplePageData struct {
	total uint        `json:"total"`
	data  interface{} `json:"data"`
}
