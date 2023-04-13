package repositories

import (
	"github.com/prometheus/common/log"
	"simple-uber/internal/models"
)

func (d *DatabaseHandler) Migrate() error {
	log.Info("running db migrations")
	err := d.pg.AutoMigrate(models.Driver{}, models.Location{}, models.Ride{}, models.Trip{}, models.User{}, models.Vehicle{})
	if err != nil {
		log.Fatalf("could not run db migrations")
	}

	log.Info("db migrations ran successfully")
	return nil
}
