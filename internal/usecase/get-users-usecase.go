package usecase

import (
	"fmt"

	"github.com/mateusgcoelho/poc-golang/internal/domain/entity"
	"github.com/mateusgcoelho/poc-golang/internal/domain/repository"
)

type GetUsersUsecase struct {
	userRepository repository.UserRepository
}

func NewGetUsersUsecase(userRepository repository.UserRepository) GetUsersUsecase {
	return GetUsersUsecase{
		userRepository: userRepository,
	}
}

func (usecase *GetUsersUsecase) Call() ([]entity.UserEntity, error) {
	users, err := usecase.userRepository.GetUsers()

	if err != nil {
		return nil, fmt.Errorf("error on search of users: %v", err);
	}

	return users, nil
}
