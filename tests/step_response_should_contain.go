package tests

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/cucumber/godog"
)

func (a *ApiFeature) TheResponseShouldContain(body *godog.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.Response.Body.Bytes(), &actual); err != nil {
		return
	}

	typeof := reflect.TypeOf(expected)

	if typeof.Kind() == reflect.Slice {
		if !issetKeyInSlice(expected, actual) {
			return fmt.Errorf("expected JSON does not contains expected keys.\n\tExpected format:\n%v\n\tActual:\n%v", expected, actual)
		}
	}

	if typeof.Kind() == reflect.Map {
		if !issetKeyInMap(expected, actual) {
			return fmt.Errorf("expected JSON does not contains expected keys.\n\tExpected format:\n%v\n\tActual:\n%v", expected, actual)
		}
	}

	return nil
}

func issetKeyInMap(expectedInterface, actualInterface interface{}) bool {
	expected, ok := expectedInterface.(map[string]interface{})
	if !ok {
		return false
	}

	actual, ok := actualInterface.(map[string]interface{})
	if !ok {
		return false
	}

	for expectedKey := range expected {
		if _, ok := actual[expectedKey]; !ok {
			return false
		}
	}

	return true
}

func issetKeyInSlice(expectedInterface, actualInterface interface{}) bool {
	expectedSlice, ok := expectedInterface.([]interface{})
	if !ok {
		return false
	}

	actualSlice, ok := actualInterface.([]interface{})
	if !ok {
		return false
	}

	if len(actualSlice) == 0 {
		return false
	}

	if len(expectedSlice) != len(actualSlice) {
		return false
	}

	actual := actualSlice[0]
	expected := expectedSlice[0]

	return issetKeyInMap(expected, actual)
}
