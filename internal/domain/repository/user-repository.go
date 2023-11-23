package repository

import (
	"github.com/mateusgcoelho/poc-golang/internal/domain/entity"
	"github.com/mateusgcoelho/poc-golang/internal/dto"
)

type UserRepository interface {
	GetUsers() ([]entity.UserEntity, error)
	CreateUser(createUserDto dto.CreateUserDto) (*entity.UserEntity, error)
}