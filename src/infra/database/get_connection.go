package database

import (
	"os"

	ssr "github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	sr "github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/repository"
	d "github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func GetConnection() d.Database {
	db := d.NewDatabase(
		d.Hostname(os.Getenv("DB_HOST")),
		d.Port(os.Getenv("DB_PORT")),
		d.User(os.Getenv("DB_USER")),
		d.Password(os.Getenv("DB_PASSWORD")),
		d.DatabaseName(os.Getenv("DB_NAME")),
		d.DbEntities(&sr.StudentModel{}, &ssr.SessionModel{}),
	)

	if err := db.OpenConnection(); err != nil {
		panic(err.Error())
	}

	return db
}
