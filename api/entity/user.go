package entity

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
