package models

type ForexData struct {
	Time string  `json:"time"`
	Date string  `json:"date"`
	Last float32 `json:"last"`
	Name string  `json:"name"`
}
