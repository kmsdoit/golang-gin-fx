package user

import (
	"go-server/lib/domain/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repository *Repository) createUser(user *user.User) error {
	return repository.db.Create(user).Error
}
