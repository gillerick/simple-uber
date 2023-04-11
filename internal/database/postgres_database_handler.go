package database

import (
	"gorm.io/gorm"
)

// PostgresDbHandler is a wrapper type for the gorm DB pointer
type PostgresDbHandler struct {
	pg *gorm.DB
}

// NewPostgresHandler sets up a new handler
func NewPostgresHandler(pg *gorm.DB) *PostgresDbHandler {
	return &PostgresDbHandler{pg: pg}
}
