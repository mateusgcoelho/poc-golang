package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/poc-golang/internal/domain/entity"
	"github.com/mateusgcoelho/poc-golang/internal/dto"
)

type UserRepositoryImpl struct {
	dbPool *pgxpool.Pool
}

func NewUserRepositoryImpl(dbPool *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		dbPool: dbPool,
	}
}

func (repository UserRepositoryImpl) GetUsers() ([]entity.UserEntity, error) {
	var users []entity.UserEntity = []entity.UserEntity{}
	
	sql := "SELECT u.id, u.name, u.email FROM users AS u"
	rows, err := repository.dbPool.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var user entity.UserEntity

		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}


func (repository UserRepositoryImpl) CreateUser(createUserDto dto.CreateUserDto) (*entity.UserEntity, error) {
	sql := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	var id *int = nil
	err := repository.dbPool.QueryRow(context.Background(), sql, createUserDto.Name, createUserDto.Email, createUserDto.Password).Scan(&id)

	if err != nil || id == nil {
		return nil, err
	}

	return nil, nil
}