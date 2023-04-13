package repositories

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"simple-uber/internal/models"
)

type UserRepository interface {
	SaveUser(user models.User) (models.User, error)
	FindUserById(id uint64) (models.User, error)
	DeleteUser(user models.User) error
}

// SaveUser creates a new user if they don't already exist in the database
func (r Repository) SaveUser(user models.User) (models.User, error) {
	result := r.db.pg.Model(models.User{}).Create(&user)
	if err := result.Error; err != nil {
		// we check if the error is a database unique constraint violation
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
			return user, errors.New("user already exists")
		}
		return models.User{}, fmt.Errorf("user could not be created %w", err)
	}
	return user, nil
}

// FindUserById searches a user by their unique ID
func (r Repository) FindUserById(userId uint64) (models.User, error) {
	var user models.User
	result := r.db.pg.Where(models.User{UserId: userId}).First(&user)
	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user does not exist")
	}
	//Handle any other error
	if err := result.Error; err != nil {
		return models.User{}, fmt.Errorf("user could not be retrieved %w", err)
	}
	return user, nil
}

// DeleteUser deletes a specified user from the database
func (r Repository) DeleteUser(user models.User) error {
	err := r.db.pg.Delete(&user).Error
	if err != nil {
		return fmt.Errorf("user could not be deleted %w", err)
	}
	return nil
}

func NewUserRepository(database *DatabaseHandler) *Repository {
	return &Repository{db: database}
}
