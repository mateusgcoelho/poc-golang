package usecase

import (
	"fmt"

	"github.com/mateusgcoelho/poc-golang/internal/domain/entity"
	"github.com/mateusgcoelho/poc-golang/internal/domain/repository"
	"github.com/mateusgcoelho/poc-golang/internal/dto"
)

type CreateUserUsecase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUsecase(userRepository repository.UserRepository) CreateUserUsecase {
	return CreateUserUsecase{
		userRepository: userRepository,
	}
}

func (usecase *CreateUserUsecase) Call(createUserDto dto.CreateUserDto) (*entity.UserEntity, error) {
	userCreated, err := usecase.userRepository.CreateUser(createUserDto)

	if err != nil {
		return nil, fmt.Errorf("error on search of users: %v", err);
	}

	return userCreated, nil
}
