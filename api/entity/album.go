package entity

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title" validate:"required"`
	Artist string  `json:"artist" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}
