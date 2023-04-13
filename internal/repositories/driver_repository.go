package repositories

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"simple-uber/internal/models"
)

type DriverRepository interface {
	SaveDriver(driver models.Driver) (models.Driver, error)
	FindDriverById(id uint64) (models.Driver, error)
	DeleteDriver(driver models.Driver) error
}

// SaveDriver creates a new driver if they don't already exist in the database
func (r Repository) SaveDriver(driver models.Driver) (models.Driver, error) {
	result := r.db.pg.Model(models.Driver{}).Create(&driver)
	if err := result.Error; err != nil {
		// we check if the error is a database unique constraint violation
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
			return driver, errors.New("driver already exists")
		}
		return models.Driver{}, fmt.Errorf("driver could not be created %w", err)
	}
	return driver, nil
}

// FindDriverById searches a driver by their unique ID
func (r Repository) FindDriverById(driverId uint64) (models.Driver, error) {
	var driver models.Driver
	result := r.db.pg.Where(models.Driver{DriverId: driverId}).First(&driver)
	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Driver{}, errors.New("driver does not exist")
	}
	//Handle any other error
	if err := result.Error; err != nil {
		return models.Driver{}, fmt.Errorf("driver could not be retrieved %w", err)
	}
	return driver, nil
}

// DeleteDriver deletes a specified driver from the database
func (r Repository) DeleteDriver(driver models.Driver) error {
	err := r.db.pg.Delete(&driver).Error
	if err != nil {
		return fmt.Errorf("driver could not be deleted %w", err)
	}
	return nil
}

func NewDriverRepository(database *DatabaseHandler) *Repository {
	return &Repository{db: database}
}
