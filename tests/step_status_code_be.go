package tests

import "fmt"

func (a *ApiFeature) TheStatusCodeShouldBe(code int) error {
	if code != a.Response.Code {
		return fmt.Errorf("expected status code %d, got %d", code, a.Response.Code)
	}

	return nil
}
