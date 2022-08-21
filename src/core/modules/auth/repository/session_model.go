package repository

import (
	"time"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"gorm.io/gorm"
)

type SessionModel struct {
	SessID       string    `gorm:"primary_key;column:session_id;type:uuid;default:uuid_generate_v4();index:idx_sess_id"`
	StudentID    string    `gorm:"column:student_id;type:uuid;unique;index:idx_student_id;not null"`
	Expiration   time.Time `gorm:"column:expiration;type:timestamp with time zone;not null"`
	AccessToken  string    `gorm:"column:access_token;type:text;not null"`
	RefreshToken string    `gorm:"column:refresh_token;type:text;not null"`
}

func (sm *SessionModel) TableName() string {
	return "sessions"
}

func (sm *SessionModel) BeforeCreate(*gorm.DB) error {
	if sm.Expiration.IsZero() {
		sm.Expiration = time.Now().UTC().Local().Add(time.Minute * 10)
	}

	return nil
}

func (sm *SessionModel) AsDomain() auth.Session {
	return auth.NewSession(
		auth.WithSessionID(sm.SessID),
		auth.WithStudentID(sm.StudentID),
		auth.WithExpiration(sm.Expiration),
		auth.WithAccessToken(sm.AccessToken),
		auth.WithRefreshToken(sm.RefreshToken),
	)
}