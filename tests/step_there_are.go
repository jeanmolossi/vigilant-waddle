package tests

import (
	"log"
	"os"
	"strings"

	"github.com/cucumber/godog"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"

	auth "github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	student "github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/repository"
)

func (a *ApiFeature) ThereAreAny(tableName string, data *godog.Table) error {
	var fields []string
	var marks []string

	head := data.Rows[0].Cells
	for _, cell := range head {
		fields = append(fields, cell.Value)
		marks = append(marks, "?")
	}

	dbConn := database.NewDatabase(
		database.Hostname(os.Getenv("DB_HOST")),
		database.Port(os.Getenv("DB_PORT")),
		database.User(os.Getenv("DB_USER")),
		database.Password(os.Getenv("DB_PASSWORD")),
		database.DatabaseName(os.Getenv("DB_NAME")),
		database.DbEntities(&student.StudentModel{}, &auth.SessionModel{}),
	)
	err := dbConn.OpenConnection()
	if err != nil {
		return err
	}

	for i := 1; i < len(data.Rows); i++ {
		var vals []interface{}
		for _, cell := range data.Rows[i].Cells {
			vals = append(vals, cell.Value)
		}

		stmt := dbConn.DB().Exec(
			`INSERT INTO `+tableName+` (`+strings.Join(fields, ",")+`) VALUES (`+strings.Join(marks, ",")+`)`,
			vals...,
		)

		if err := stmt.Error; err != nil {
			log.Println("Statement err", err)
			return err
		}
	}

	return nil
}

func (a *ApiFeature) ClearDB(*godog.Scenario) error {
	dbConn := database.NewDatabase(
		database.Hostname(os.Getenv("DB_HOST")),
		database.Port(os.Getenv("DB_PORT")),
		database.User(os.Getenv("DB_USER")),
		database.Password(os.Getenv("DB_PASSWORD")),
		database.DatabaseName(os.Getenv("DB_NAME")),
	)
	err := dbConn.OpenConnection()
	if err != nil {
		return err
	}

	tables := []string{
		"sessions",
		"users",
	}

	for _, table := range tables {
		stmt := dbConn.DB().Exec(`DELETE FROM ` + table)
		if err := stmt.Error; err != nil {
			return err
		}
	}

	return nil
}
