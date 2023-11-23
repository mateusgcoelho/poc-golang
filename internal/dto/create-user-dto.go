package dto

import "fmt"

type CreateUserDto struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (u *CreateUserDto) Validate() error {
	if (u.Name == "" && u.Email == "" && u.Password == "") {
		return fmt.Errorf("request body is empty or malformed")
	}
	if u.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if u.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if u.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}