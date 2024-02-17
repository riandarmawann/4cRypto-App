package controller

import (
	"net/http"
	"strings"

	"4crypto/config"
	"4crypto/model/dto"
	"4crypto/usecase"
	"4crypto/utils/common"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc         usecase.AuthUseCase
	rg         *gin.RouterGroup
	jwtService common.JwtToken
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwtService common.JwtToken) *AuthController {
	return &AuthController{uc: uc, rg: rg, jwtService: jwtService}
}

func (c *AuthController) Route() {
	authGroup := c.rg.Group(config.AuthGroup)
	authGroup.POST(config.AuthLogin, c.loginHandler)
	authGroup.GET(config.AuthRefreshToken, c.refreshTokenHandler)
}

func (c *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := c.uc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Successfully logged in", response)
}

func (c *AuthController) refreshTokenHandler(ctx *gin.Context) {
	tokenString := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	newToken, err := c.jwtService.RefreshToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	common.SendCreateResponse(ctx, "token has been refreshed successfully", newToken)
}
