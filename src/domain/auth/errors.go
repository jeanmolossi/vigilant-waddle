package auth

import "errors"

var ErrHasNotSession = errors.New("has not session")
var ErrHasNotSessionID = errors.New("has not session identifier")
var ErrHasNotStudentID = errors.New("has not student identifier")
