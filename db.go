package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// a postgres db
type Postgres struct {
	Host, User, Password, DbName, Port string
	SslModeEnable                      bool
}

// return p's DSN (Data Source Name)
func (p Postgres) DSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=enable", p.Host, p.User, p.Password, p.DbName, p.Port)
	if !p.SslModeEnable {
		dsn += " sslmode=disable"
	}
	return dsn
}

// establish a db (postgres) connection and return a pointer to a gorm.DB instance.
// return a non-nil error if something went wrong
func (p Postgres) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(p.DSN()), &gorm.Config{})
}
