package repository

import (
	"time"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/auth"
	"gorm.io/gorm"
)

// SessionModel represents how a session is stored in the database
type SessionModel struct {
	SessID       string    `gorm:"primary_key;column:session_id;type:uuid;default:uuid_generate_v4();index:idx_sess_id"`
	UserID       string    `gorm:"column:usr_id;type:uuid;unique;index:idx_usr_id;not null"`
	Expiration   time.Time `gorm:"column:expiration;type:timestamp with time zone;not null"`
	AccessToken  string    `gorm:"column:access_token;type:text;not null"`
	RefreshToken string    `gorm:"column:refresh_token;type:text;not null"`
}

// TableName overrides the table name used by SessionModel to `sessions`
//
// Read more about GORM conventions:
//
// https://gorm.io/docs/conventions.html#TableName
func (sm *SessionModel) TableName() string {
	return "sessions"
}

// BeforeCreate is a hook to set the expiration field
//
// Read more about GORM hooks:
//
// https://gorm.io/docs/hooks.html#Creating-an-object
func (sm *SessionModel) BeforeCreate(*gorm.DB) error {
	if sm.Expiration.IsZero() {
		sm.Expiration = time.Now().UTC().Local().Add(time.Minute * 10)
	}

	return nil
}

// AsDomain will return a SessionModel as auth.Session
func (sm *SessionModel) AsDomain() auth.Session {
	return auth.NewSession(
		auth.WithSessionID(sm.SessID),
		auth.WithUserID(sm.UserID),
		auth.WithExpiration(sm.Expiration),
		auth.WithAccessToken(sm.AccessToken),
		auth.WithRefreshToken(sm.RefreshToken),
	)
}
