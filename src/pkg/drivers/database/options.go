package database

import (
	"gorm.io/gorm"
)

// Hostname will set option host for database connection
func Hostname(hostname string) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.host = hostname
			return nil
		}

		return ErrCanParseDbInstance
	}
}

// Port will set option port for database connection
func Port(port string) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.port = port
			return nil
		}

		return ErrCanParseDbInstance
	}
}

// User will set option user for database connection
func User(user string) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.user = user
			return nil
		}

		return ErrCanParseDbInstance
	}
}

// Password will set option password for database connection
func Password(pass string) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.pass = pass
			return nil
		}

		return ErrCanParseDbInstance
	}
}

// DatabaseName will set option dbname for database connection
func DatabaseName(dbname string) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.dbname = dbname
			return nil
		}

		return ErrCanParseDbInstance
	}
}

// Dialect will set option dialect for database connection
func Dialect(dialect func(dsn string) gorm.Dialector) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.dialect = dialect
			return nil
		}

		return ErrCanParseDbInstance
	}
}

// DbEntities will set option entities for database connection
//
// NOTE: this option is only for auto migrate.
//
// That will set the entities that will be auto migrated.
func DbEntities(dst ...interface{}) DbOption {
	return func(d Database) error {
		if db, ok := d.(*database); ok {
			db.autoMigrateEntities = dst
			return nil
		}

		return ErrCanParseDbInstance
	}
}
