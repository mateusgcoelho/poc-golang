package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	data "github.com/mateusgcoelho/poc-golang/internal/data/repository"
	"github.com/mateusgcoelho/poc-golang/internal/domain/repository"
	"github.com/mateusgcoelho/poc-golang/internal/dto"
	"github.com/mateusgcoelho/poc-golang/internal/usecase"
)

var (
	getUsersUsecase usecase.GetUsersUsecase
	createUserUsecase usecase.CreateUserUsecase
)

func injectUserHandlerDependencies() {
	userRepository := repository.UserRepository(data.NewUserRepositoryImpl(dbPool))

	getUsersUsecase = usecase.NewGetUsersUsecase(userRepository)
	createUserUsecase = usecase.NewCreateUserUsecase(userRepository)
}

func GetUsers(ctx *gin.Context) {
	users, err := getUsersUsecase.Call()

	if err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "An error occurred when searching for users", err)
		return 
	}

	sendSuccessResponse(ctx, http.StatusOK, "get-users", users)
}

func CreateUser(ctx *gin.Context) {
	request  := dto.CreateUserDto{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	createUserDto := dto.CreateUserDto{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}

	userCreated, err := createUserUsecase.Call(createUserDto)

	if err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "An error occurred when create a new user", err)
		return 
	}

	sendSuccessResponse(ctx, http.StatusOK, "create-user", userCreated)
}