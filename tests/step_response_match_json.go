package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cucumber/godog"
	"github.com/r3labs/diff"
)

func (a *ApiFeature) TheResponseMatchJSON(body *godog.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.Response.Body.Bytes(), &actual); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		changelog, _ := diff.Diff(expected, actual)
		log.Printf("CHANGE LOG\n%+v", changelog)
		return fmt.Errorf("expected JSON does not match actual.\n\tExpected:\n%v\n\tActual:\n%v", expected, actual)
	}

	return nil
}
