package repositories

import "simple-uber/internal/models"

type DriverRepository interface {
	SaveDriver(driver models.Driver) (*models.Driver, error)
	FindDriverById(id uint64) (*models.Driver, error)
	DeleteDriver(driver models.Driver) error
}
