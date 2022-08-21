package student

import "errors"

var ErrInvalidStudentProp = errors.New("invalid student property")
var ErrNoDataToSync = errors.New("no data to sync")
var ErrEmptyPropertyValue = errors.New("empty property value")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrNoStudentsFound = errors.New("no students found")
var ErrInvalidCredentials = errors.New("invalid credentials")

// GetMeOptions validations
var ErrMissingStudentID = errors.New("missing student id")
var ErrStudentIDInvalid = errors.New("student id invalid")
