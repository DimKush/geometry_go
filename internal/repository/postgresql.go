package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	warehouses = "warehouses"
	units = "units"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Timezone string
	SSLMode  string
}

func NewPostgresConnection(cfg Config) (*gorm.DB, error) {
	var DSN string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Dbname,
		cfg.Port,
		cfg.SSLMode,
		cfg.Timezone)

	fmt.Printf("\nDSN : %s", DSN)
	dialector := postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	})

	return gorm.Open(dialector, &gorm.Config{})
}
