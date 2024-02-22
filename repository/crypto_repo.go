package repository

import "database/sql"

type CryptoRepository interface {
	// GetById(id string) (entity.User, error)
	// GetByUsername(username string) (entity.User, error)
	// DeleteUser(id string) error
}

type cryptoRepository struct {
	db *sql.DB
}

func NewCryptoRepository(db *sql.DB) CryptoRepository {
	return &cryptoRepository{db: db}
}

// func (cr)
