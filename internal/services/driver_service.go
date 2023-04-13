package services

import (
	"fmt"
	"simple-uber/internal/models"
	"simple-uber/internal/repositories"
)

type DriverService struct {
	repository repositories.DriverRepository
}

type SimpleUberDriverService interface {
	SaveDriver(models.Driver) (models.Driver, error)
	FindDriverByUserId(driverId uint64) (models.Driver, error)
	DeleteDriver(driver models.Driver) error
}

func (s DriverService) SaveDriver(driver models.Driver) (models.Driver, error) {
	savedDriver, err := s.repository.SaveDriver(driver)
	if err != nil {
		return models.Driver{}, fmt.Errorf("driver creation failed with error: %w", err)
	}
	return savedDriver, nil
}

func (s DriverService) FindDriverByUserId(driverId uint64) (models.Driver, error) {
	driver, err := s.repository.FindDriverById(driverId)
	if err != nil {
		return models.Driver{}, fmt.Errorf("account retrieval failed with error: %w", err)
	}
	return *driver, nil
}

func (s DriverService) DeleteDriver(driver models.Driver) error {
	err := s.repository.DeleteDriver(driver)
	if err != nil {
		return fmt.Errorf("driver deletion failed with error: %w", err)
	}
	return nil
}

func NewDriverService(repository repositories.DriverRepository) *DriverService {
	return &DriverService{repository: repository}
}
