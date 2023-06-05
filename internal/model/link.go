package model

type Link struct {
	Link string `json:"link,omitempty"`
}

type LinkResponse struct {
	Link
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
