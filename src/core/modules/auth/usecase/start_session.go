package usecase

import (
	"context"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/filters"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/paginator"
)

func StartSession(
	ctx context.Context,
	repo auth.CreateSessionRepository,
	usrRepo student.GetStudentsRepository,
) auth.StartSession {
	getStudent := func(username, password string) (student.Student, error) {
		f := filters.NewConditions()
		f.WithCondition("usr_email", username)

		students, err := usrRepo.Run(ctx, f, paginator.New())
		if err != nil {
			return nil, err
		}

		// if no students were found, return an error
		// when it happens no user with that email exists or
		// more than 1 user with that email exists
		if len(students) != 1 {
			return nil, student.ErrInvalidCredentials
		}

		// get the student found
		s := students[0]
		// validate user password to return the student
		if !s.IsPasswordValid(password) {
			return nil, student.ErrInvalidCredentials
		}

		return s, nil
	}

	makeSession := func(student student.Student) (auth.Session, error) {
		// creates new access token and new refresh token
		accessToken := auth.NewSessionToken(student.GetID())
		refreshToken := auth.NewSessionToken(
			student.GetID(),
			auth.WithIssuer("refresh_provider"),
			auth.ExpiresIn(auth.RefreshExpiration),
		)

		// create new session with the student id, access token and refresh token
		s := auth.NewSession(
			auth.WithUserID(student.GetID()),
			auth.WithAccessToken(accessToken.Token()),
			auth.WithRefreshToken(refreshToken.Token()),
			auth.WithExpiration(accessToken.Expiration()),
		)

		newSession, err := repo.Run(ctx, s)
		if err != nil {
			return nil, err
		}

		return newSession, nil
	}

	return func(username, password string) (auth.Session, error) {
		student, err := getStudent(username, password)
		if err != nil {
			return nil, err
		}

		return makeSession(student)
	}
}
