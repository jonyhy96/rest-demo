package models

// PageResponse response paged response
type PageResponse struct {
	Total    string      `json:"total"`
	Page     string      `json:"page"`
	PageSize string      `json:"pageSize"`
	Data     interface{} `json:"data"`
}
