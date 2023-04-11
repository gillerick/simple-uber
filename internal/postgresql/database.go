package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple-uber/internal/configs"
	// Add postgresql support
	_ "github.com/lib/pq"
)

func NewConnection(config configs.Database) (*gorm.DB, error) {
	var err error
	dbDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not create a connection: %w", err)
	}
	return db, nil
}
