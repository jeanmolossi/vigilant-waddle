package producer

import "errors"

var ErrInvalidProducerProp = errors.New("invalid producer property")
var ErrNoDataToSync = errors.New("no data to sync")
var ErrEmptyPropertyValue = errors.New("empty property value")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrNoProducersFound = errors.New("no producers found")
var ErrInvalidCredentials = errors.New("invalid credentials")

// GetMeOptions validations
var ErrMissingProducerID = errors.New("missing producer id")
var ErrProducerIDInvalid = errors.New("producer id invalid")
