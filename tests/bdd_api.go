package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

type ApiFeature struct {
	BaseURL  string
	Response *httptest.ResponseRecorder
	Token    string
}

func (a *ApiFeature) MakeRequestAuth(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.Token)
}

func (a *ApiFeature) ResetResponse(*godog.Scenario) {
	a.Response = httptest.NewRecorder()
}

func (a *ApiFeature) MakeRequest(req *http.Request) (err error) {
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	a.Response.WriteHeader(res.StatusCode)
	if err != nil {
		fmt.Fprintf(a.Response, "Error: %s", err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(a.Response, "Error: %s", err)
		return nil
	}

	a.Response.Header().Set("Content-Type", "application/json")
	a.Response.WriteHeader(res.StatusCode)
	fmt.Fprintf(a.Response, "%s", string(data))
	return nil
}
