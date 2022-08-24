package tests

import (
	"context"
	"fmt"
	"os"

	auth "github.com/jeanmolossi/vigilant-waddle/src/core/modules/auth/repository"
	student "github.com/jeanmolossi/vigilant-waddle/src/core/modules/student/repository"
	authDmn "github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

func (a *ApiFeature) ThereIsUserLogged(usrID string) error {
	dbConn := database.NewDatabase(
		database.Hostname(os.Getenv("DB_HOST")),
		database.Port(os.Getenv("DB_PORT")),
		database.User(os.Getenv("DB_USER")),
		database.Password(os.Getenv("DB_PASSWORD")),
		database.DatabaseName(os.Getenv("DB_NAME")),
		database.DbEntities(&student.StudentModel{}, &auth.SessionModel{}),
	)
	err := dbConn.OpenConnection()
	if err != nil {
		return err
	}

	usr := &student.StudentModel{
		ID: usrID,
	}

	result := dbConn.DB().First(usr)
	if result.Error != nil {
		return fmt.Errorf("err during get usr: %v", result.Error.Error())
	}

	accessToken := authDmn.NewSessionToken(usrID)
	refreshToken := authDmn.NewSessionToken(
		usrID,
		authDmn.WithIssuer("refresh_provider"),
		authDmn.ExpiresIn(authDmn.RefreshExpiration),
	)

	// create new session with the student id, access token and refresh token
	s := authDmn.NewSession(
		authDmn.WithStudentID(usrID),
		authDmn.WithAccessToken(accessToken.Token()),
		authDmn.WithRefreshToken(refreshToken.Token()),
		authDmn.WithExpiration(accessToken.Expiration()),
	)

	newSession, err := auth.NewCreateSession(dbConn).Run(
		context.Background(), s)
	if err != nil {
		return err
	}

	a.Token = newSession.HashToken()

	return nil
}
