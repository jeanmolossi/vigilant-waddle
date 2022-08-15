package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database is the Database interface implementation.
//
// That struct stores the gorm.DB instance, and the dsn configurations
type database struct {
	// db is the gorm.DB instance.
	// It is a singleton.
	db *gorm.DB
	// once will do a singleton connection to the database.
	once *sync.Once

	// Following are the database configurations.

	host   string
	port   string
	user   string
	pass   string
	dbname string

	// dialect is the database dialect.
	// It is a function that returns a gorm.Dialector.
	// That dialector can be used to set up different database dialects.
	// You can use MySQL, PostgreSQL, SQLite, etc.
	dialect func(dsn string) gorm.Dialector

	// autoMigrateEntities is a list of entities that will be auto migrated.
	autoMigrateEntities []interface{}
}

func NewDatabase(dbOptions ...DbOption) Database {
	d := &database{
		once:                &sync.Once{},
		host:                "localhost",
		port:                "5432",
		dialect:             postgres.Open,
		autoMigrateEntities: make([]interface{}, 0),
	}

	if len(dbOptions) > 0 {
		for _, opt := range dbOptions {
			if err := opt(d); err != nil {
				panic(err.Error())
			}
		}
	}

	return d
}

// OpenConnection will make a connection to the database.
//
// The package database implementations makes a singleton to the connection.
func (db *database) OpenConnection() (err error) {
	db.once.Do(func() {
		db.db, err = db.connect()
	})

	return
}

// DB returns the gorm.DB instance.
func (db *database) DB() *gorm.DB {
	return db.db
}

// connect is the callback that will make a connection to the database.
func (db *database) connect() (*gorm.DB, error) {
	// level is Info by default.
	level := logger.Info

	// when production environment we log only errors.
	if os.Getenv("ENVIRONMENT") == "production" {
		level = logger.Error
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: level,
			Colorful: true,
		},
	)
	// open a connection to the database.
	dbConnection, err := gorm.Open(
		db.dialect(db.dsn()),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		return nil, err
	}

	// set up the database connection.
	db.db = dbConnection

	// auto migrate the entities.
	db.autoMigrate()

	return nil, nil
}

// dsn returns the database connection string.
func (db *database) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		db.host, db.port, db.user, db.dbname, db.pass,
	)
}

// autoMigrate will auto migrate the entities.
func (db *database) autoMigrate() error {
	if os.Getenv("ENVIRONMENT") == "production" {
		return nil
	}

	if len(db.autoMigrateEntities) == 0 {
		return nil
	}

	if err := db.db.AutoMigrate(db.autoMigrateEntities...); err != nil {
		return err
	}

	return nil
}
