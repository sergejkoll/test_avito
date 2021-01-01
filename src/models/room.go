package models

type Room struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	DateAdded   int64  `json:"date_added"`
}
