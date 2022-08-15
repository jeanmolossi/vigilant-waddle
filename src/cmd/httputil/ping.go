// Package httputil represents the http responses documentation.
package httputil

// PinkOk example
type PingOk struct {
	Message string `json:"message" example:"pong"`
}

// PingInternalServerErr example
type PingInternalServerErr struct {
	Message string `json:"message" example:"unexpected error"`
}
