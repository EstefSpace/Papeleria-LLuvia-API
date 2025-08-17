package models

/*
Structs de ventas
*/

type Sale struct {
	ID       *string   `json:"id,omitempty"`
	User     string    `json:"user"`
	Total    float64   `json:"total"`
	Date     string    `json:"date"`
	Products []Product `json:"products"`
}
