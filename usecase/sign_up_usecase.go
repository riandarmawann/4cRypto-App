package usecase

import (
	"4crypto/model/entity"
	"4crypto/repository"
	"errors"
	"fmt"
)

type SignUpUseCase interface {
	SignUp(newUser entity.User) error
}

type signUpUseCase struct {
	userRepo repository.UserRepository
}

func NewSignUpUseCase(userRepo repository.UserRepository) SignUpUseCase {
	return &signUpUseCase{userRepo: userRepo}
}

func (su *signUpUseCase) SignUp(newUser entity.User) error {
	// Check if username or email already exists
	_, err := su.userRepo.GetByUsername(newUser.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	// _, err = su.userRepo.GetByUsername(newUser.Email)
	// if err == nil {
	// 	return errors.New("email already exists")
	// }

	// Add the new user
	err = su.userRepo.Create(newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
