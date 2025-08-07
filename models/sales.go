package models

/*
Structs de ventas
*/

type Sale struct {
	ID       *string   `json:"id,omitempty"`
	User     string    `json:"user"`
	Total    int       `json:"total"`
	Date     string    `json:"date"`
	Products []Product `json:"products"`
}
