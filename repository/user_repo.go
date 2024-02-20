package repository

import (
	"4crypto/utils/common/query"
	"database/sql"
	"fmt"
	"time"

	"4crypto/model/entity"
)

type UserRepository interface {
	Create(user entity.User) error
	GetById(id string) (entity.User, error)
	GetByUsername(username string) (entity.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, updatedUser entity.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) error {
	query := `INSERT INTO users (id, name, email, username, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW());`

	fmt.Println(query)

	_, err := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Username,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	)

	user.Password = ""

	if err != nil {
		return err
	}

	return nil

}

func (r *userRepository) GetById(id string) (entity.User, error) {
	var user entity.User
	query := `SELECT id, name, email, username, password, role, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetByUsername(username string) (entity.User, error) {
	var user entity.User

	err := r.db.QueryRow(query.GetUserByUsername, username).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(id string, updatedUser entity.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, username = $3, password = $4, role = $5, updated_at = NOW()
		WHERE id = $6
	`
	_, err := r.db.Exec(query,
		updatedUser.Name,
		updatedUser.Email,
		updatedUser.Username,
		updatedUser.Password,
		updatedUser.Role,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
