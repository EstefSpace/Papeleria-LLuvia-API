package models

/*
Esto es una estructura, como bien lo dice "struct"
Utilizo modelos aqui para poder buscar información en una base de datos (en este caso MongoDB)
*/
type Product struct {
	ID     *string `json:"id,omitempty"`
	Name   string  `json:"name"` // Aca le ponemos nombre a la key, luego el tipo de dato (string) y luego como apareceria el nombre en un .json
	Amount int     `json:"amount"`
	Price  int     `json:"price"`
}

type DeleteProduct struct {
	ID string `json:"id"`
}
