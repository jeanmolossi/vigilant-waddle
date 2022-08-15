package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func (a *ApiFeature) DoRegister() error {
	registerEndpoint := fmt.Sprintf("%s%s", a.BaseURL, "/students/register")

	resp, err := http.Post(registerEndpoint, "application/json",
		strings.NewReader(`{"password": "123456789","username": "john@doe.com"}`),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		bytes, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(bytes), "already exists") {
			log.Println(string(bytes))
			return errors.New("failed to register student")
		}
	}

	return nil
}

func (a *ApiFeature) DoLogin() error {
	loginEndpoint := fmt.Sprintf("%s%s", a.BaseURL, "/auth/login")

	resp, err := http.Post(loginEndpoint, "application/json",
		strings.NewReader(`{"password": "123456789","username": "john@doe.com"}`),
	)
	if err != nil {
		return err
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type body struct {
		AccessToken string `json:"access_token"`
	}

	var b body
	err = json.Unmarshal(response, &b)
	if err != nil {
		return err
	}

	a.Token = b.AccessToken

	return nil
}
