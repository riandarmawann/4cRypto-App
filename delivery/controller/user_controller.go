package controller

import (
	"4crypto/model/dto/res"
	"4crypto/model/entity"
	"4crypto/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
	rg          *gin.RouterGroup
}

func NewUserController(userUseCase usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		rg:          rg,
	}
}

func (u *UserController) FindById(ctx *gin.Context) {
	userID := ctx.Param("id")

	var res res.CommonResponse

	user, err := u.userUseCase.FindById(userID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Retrieved data successfully"
	res.Data = user

	ctx.JSON(http.StatusOK, res)
}

func (u *UserController) Create(ctx *gin.Context) {

	var user entity.User

	ctx.BindJSON(&user)

	var res res.CommonResponse

	err := u.userUseCase.Create(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusCreated
	res.Status = "Success"
	res.Message = "Create data successfully"

	ctx.JSON(res.Code, res)
}
