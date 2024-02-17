package common

import (
	"errors"
	"log"
	"time"

	"4crypto/config"
	"4crypto/model/dto"
	"4crypto/model/entity"
	"4crypto/utils/model_util"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(payload entity.User) (dto.AuthResponseDto, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	RefreshToken(oldTokenString string) (dto.AuthResponseDto, error)
}

type jwtToken struct {
	cfg config.TokenConfig
}

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg}
}

func (j *jwtToken) GenerateToken(payload entity.User) (dto.AuthResponseDto, error) {
	claims := model_util.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtLifeTime)),
		},
		UserId: payload.Id,
		Role:   payload.Role,
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("failed to generate token")
	}

	return dto.AuthResponseDto{Token: token}, nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to verify token")
	}

	// convert dari token.Claims ke jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	log.Println("INI BANG >>>>>>>>>>>>>>>>>>>>>>>>>>>> ", claims["exp"])
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return nil, errors.New("invalid claim token")
	}

	return claims, nil
}

func (j *jwtToken) RefreshToken(refreshToken string) (dto.AuthResponseDto, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return dto.AuthResponseDto{}, errors.New("invalid claim token")
	}

	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: newTokenString}, nil
}
