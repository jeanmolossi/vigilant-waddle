package student

import "errors"

var ErrInvalidStudentProp = errors.New("invalid student property")
var ErrNoDataToSync = errors.New("no data to sync")
var ErrEmptyPropertyValue = errors.New("empty property value")
var ErrEmailAlreadyExists = errors.New("email already exists")
