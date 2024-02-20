package controller

import (
	"4crypto/config"
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

func (c *UserController) Route() {
	userGroup := c.rg.Group(config.UserGroup)
	//userGroup.POST(config.CreateUser, c.CreateUser)
	//userGroup.GET(config.UserGetByID, c.GetUserByID)

	userGroup.DELETE(config.DeleteUserByID, c.DeleteUserByID)
	userGroup.PUT(config.UpdateUserByID, c.UpdateUserByID)
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

func (uc *UserController) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.userUseCase.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (uc *UserController) UpdateUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedUser entity.User
	if err := ctx.BindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := uc.userUseCase.UpdateUser(id, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
