package usecase

import (
	"errors"
	"fmt"

	"4crypto/model/entity"
	"4crypto/repository"
)

type UserUseCase interface {
	Create(user entity.User) error
	FindById(id string) (entity.User, error)
	FindByUsernamePassword(username string, password string) (entity.User, error)
	DeleteById(id string) error
	UpdateUser(id string, newUser entity.User) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) Create(user entity.User) error {
	err := u.userRepo.Create(user)
	if err != nil {
		return fmt.Errorf("failed to create data customer")
	}
	return err
}

func (u *userUseCase) FindById(id string) (entity.User, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) FindByUsernamePassword(username string, password string) (entity.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return entity.User{}, errors.New("invalid username or password")
	}

	// compare password
	if user.Password != password {
		return entity.User{}, err
	}

	// set user password jadi kosong agar tidak ditampilkan di response
	user.Password = ""
	return user, nil
}

func (u *userUseCase) DeleteById(id string) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %s: %v", id, err)
	}
	return nil
}

func (u *userUseCase) UpdateUser(id string, newUser entity.User) error {
	// Mendapatkan user yang ingin diupdate
	existingUser, err := u.userRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("user with id %s not found", id)
	}

	// Update field yang diperlukan
	existingUser.Username = newUser.Username
	existingUser.Password = newUser.Password

	// Simpan perubahan ke repository
	err = u.userRepo.UpdateUser(id, existingUser)
	if err != nil {
		return fmt.Errorf("failed to update user with ID %s: %v", id, err)
	}
	return nil
}
