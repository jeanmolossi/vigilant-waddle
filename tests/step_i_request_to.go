package tests

import (
	"fmt"
	"net/http"
)

func (a *ApiFeature) IRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method,
		fmt.Sprintf("%s%s", a.BaseURL, endpoint),
		nil)
	if err != nil {
		return err
	}

	a.MakeRequestAuth(req)
	err = a.MakeRequest(req)

	return
}
