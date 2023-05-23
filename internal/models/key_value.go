package models

type KeyValueIncrementRequest struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type KeyValueIncrementResponse struct {
	Value int `json:"value"`
}
