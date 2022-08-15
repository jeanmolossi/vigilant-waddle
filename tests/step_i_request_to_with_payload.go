package tests

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
)

func (a *ApiFeature) IRequestToWithPayload(method, endpoint string, payload *godog.DocString) (err error) {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", a.BaseURL, endpoint),
		strings.NewReader(payload.Content))
	if err != nil {
		return
	}

	a.MakeRequestAuth(req)
	err = a.MakeRequest(req)
	return
}
