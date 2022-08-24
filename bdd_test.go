package main

import (
	"context"
	"net/http/httptest"
	"os"

	"github.com/cucumber/godog"
	"github.com/jeanmolossi/vigilant-waddle/tests"
)

type apiFeature struct {
	baseURL string
	resp    *httptest.ResponseRecorder
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &tests.ApiFeature{
		BaseURL: os.Getenv("API_HOST"),
		Token:   "",
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.ResetResponse(sc)
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		api.ClearDB(sc)
		return ctx, nil
	})

	ctx.Step(`^I "(GET|POST|PUT|DELETE)" to "([^"]*)"$`, api.IRequestTo)
	ctx.Step(`^I "(POST|PUT)" to "([^"]*)" with:$`, api.IRequestToWithPayload)
	ctx.Step(`^the status code received should be (\d+)$`, api.TheStatusCodeShouldBe)
	ctx.Step(`^the response received should match json:$`, api.TheResponseMatchJSON)
	ctx.Step(`^the response should contain:$`, api.TheResponseShouldContain)
	ctx.Step(`^there is user "([^"]*)" logged`, api.ThereIsUserLogged)
	ctx.Step(`^there are headers:$`, api.ThereAreHeaders)
	ctx.Step(`^there are "([^"]*)" with:$`, api.ThereAreAny)
}
