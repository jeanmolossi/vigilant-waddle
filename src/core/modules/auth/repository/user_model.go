package repository

import (
	"time"
)

// UserModel represents how a student is stored in the database.
type UserModel struct {
	ID        string    `gorm:"primary_key;column:usr_id;type:uuid;default:uuid_generate_v4();index:idx_usr_id"`
	Email     string    `gorm:"column:usr_email;unique;type:varchar(255);"`
	Type      string    `gorm:"column:usr_scope;type:varchar(64);default:student"`
	CreatedAt time.Time `gorm:"column:usr_created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt time.Time `gorm:"column:usr_updated_at;type:timestamp with time zone;default:now()"`
}

// TableName overrides the table name used by UserModel to `users`
//
// Read more about GORM conventions:
//
// https://gorm.io/docs/conventions.html#TableName
func (UserModel) TableName() string {
	return "users"
}
