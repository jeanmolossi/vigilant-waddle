package tests

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/cucumber/godog"
	"github.com/jeanmolossi/vigilant-waddle/src/infra/database"
)

func (a *ApiFeature) ThereAreHeaders(data *godog.Table) {
	if a.Headers == nil {
		a.Headers = make(map[string]string)
	}

	for _, row := range data.Rows {
		if len(row.Cells) < 2 {
			continue
		}

		name := row.Cells[0].Value
		value := row.Cells[1].Value

		a.Headers[name] = compileValue(value)
	}
}

// transformWithMethod will return the method name who will call
// the method to transform a value
func transformWithMethod(value string) string {
	value = strings.TrimSpace(value)
	// if the value received contains starting string:
	// @transform:exampleMethod:...
	//
	// It means: The value will be transformed with
	// exampleMethod
	pattern := regexp.MustCompile(`@transform:(?P<method>(\w+)):.+`)
	if !pattern.MatchString(value) {
		return ""
	}

	groups := pattern.SubexpNames()
	for _, match := range pattern.FindAllStringSubmatch(value, -1) {
		for i, group := range match {
			if groups[i] == "method" {
				return group
			}
		}
	}

	return ""
}

// dbSeachOptions will contains the options
// who will be used to search a value from db
type dbSearchOptions struct {
	tableName string
	column    string
	value     string
}

// isValueFromDb will get value and check if it is a query
// in database. If it is a db query will return the
// dbSearchOptions pointer, else will be nil
func isValueFromDb(value string) *dbSearchOptions {
	value = strings.TrimSpace(value)

	// checks if value matches with pattern who indicates
	// a db query:
	// If value contains:
	//	@db:tableName:columnName:columnValue
	//
	// It means the value should be the value returned from
	// the following query:
	//	SELECT * FROM tableName WHERE columnName = columnValue
	pattern := regexp.MustCompile(`@db:(?P<table>(\w+)):(?P<column>(.+)):(?P<value>(.{0,255}))$`)

	if !pattern.MatchString(value) {
		return nil
	}

	groups := pattern.SubexpNames()

	searchOpts := &dbSearchOptions{}
	for _, match := range pattern.FindAllStringSubmatch(value, -1) {
		for i, group := range match {
			switch groups[i] {
			case "table":
				searchOpts.tableName = group
			case "column":
				searchOpts.column = group
			case "value":
				searchOpts.value = group
			}
		}
	}

	return searchOpts
}

// searchFromDB will receive a dbSearchOptions pointer
// and will do a db query like:
//	SELECT * FROM opt.tableName WHERE opt.columnName = opt.columnValue
func searchFromDB(opt *dbSearchOptions) interface{} {
	if opt == nil {
		return nil
	}

	db := database.GetConnection()
	db.OpenConnection()

	search := db.DB().Table(opt.tableName)
	statement := fmt.Sprintf("%s = ?", opt.column)
	result := make(map[string]interface{}, 0)
	search.Where(statement, opt.value).Limit(1).Find(&result)

	return result
}

// compileValue will checks if the value received should pass
// by a compilation value like transform or db query
func compileValue(value string) string {
	// if value is empty still returning empty
	if value == "" {
		return ""
	}

	var dbValue interface{}
	// checks if is a value who needs a db query
	if valueFromDB := isValueFromDb(value); valueFromDB != nil {
		// query the value in database
		dbValue = searchFromDB(valueFromDB)
	}

	var transformedValue interface{}
	// checks if the value received also needs a method transformation
	if methodName := transformWithMethod(value); methodName != "" {
		// when needs a method transformation, transform the received value
		// if dbValue is not nil transform value with that.
		if dbValue != nil {
			transformedValue = methods(methodName, dbValue)
		} else {
			// when dbValue is nill transform original value
			transformedValue = methods(methodName, value)
		}
	}

	if transformedValue != nil {
		return transformedValue.(string)
	}

	return value
}

// methods will receive a key and a value to transform
// the key corresponds to a method case
func methods(key string, value interface{}) interface{} {
	switch key {
	case "toAccessToken":
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return nil
		}

		studentID := valueMap["student_id"]
		sessionID := valueMap["session_id"]
		return base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s", studentID, sessionID)),
		)

	default:
		return nil
	}
}
