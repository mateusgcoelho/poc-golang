package entity

type UserEntity struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
}