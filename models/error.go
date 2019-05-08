package models

// Error is struct for error
type Error struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// HTTPError is struct for HTTPError
type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
