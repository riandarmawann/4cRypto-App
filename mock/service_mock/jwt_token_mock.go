package servicemock

import (
	"4crypto/model/dto"
	"4crypto/model/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type JwtTokenMock struct {
	mock.Mock
}

func (j *JwtTokenMock) GenerateToken(payload entity.User) (dto.AuthResponseDto, error) {
	args := j.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (j *JwtTokenMock) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	args := j.Called(tokenString)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

func (j *JwtTokenMock) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	args := j.Called(oldTokenString)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}
