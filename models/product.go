package models

type Product struct {
	Id uint              `json:"id"`
	Title string				 `json:"title"`
	Desctiption string	 `json:"description"`
	Image string				 `json:"image"`
	Price float64				 `json:"price"`
}