// Package database contains the database drivers.
//
// Actually, this package implements the DB interface from https://gorm.io
package database

import (
	"gorm.io/gorm"
)

// DbOption is a function that sets a database option.
type DbOption func(d Database) error

// Database is an abstraction who wraps the gorm.DB.
//
// That wrapper target to extract maximum of code from the gorm.DB.
type Database interface {
	// OpenConnection will make a connection to the database.
	//
	// The package database implementations makes a singleton to the connection.
	OpenConnection() error
	// DB returns the gorm.DB instance.
	DB() *gorm.DB
	// dsn returns the database connection string.
	dsn() string
}
