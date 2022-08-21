package repository

import (
	"time"

	"gorm.io/gorm"
)

var studentsTable = StudentModel{}.TableName()

// StudentModel represents how a student is stored in the database.
type StudentModel struct {
	ID        string    `gorm:"primary_key;column:usr_id;type:uuid;default:uuid_generate_v4();index:idx_usr_id"`
	Email     string    `gorm:"column:usr_email;unique;type:varchar(255);"`
	Password  string    `gorm:"column:usr_password;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:usr_created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt time.Time `gorm:"column:usr_updated_at;type:timestamp with time zone;default:now()"`
}

// TableName overrides the table name used by StudentModel to `users`
//
// Read more about GORM conventions:
//
// https://gorm.io/docs/conventions.html#TableName
func (StudentModel) TableName() string {
	return "users"
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
//
// Read more about GORM hooks:
//
// https://gorm.io/docs/hooks.html#Creating-an-object
func (s *StudentModel) BeforeCreate(*gorm.DB) error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
//
// Read more about GORM hooks:
//
// https://gorm.io/docs/hooks.html#Updating-an-object
func (s *StudentModel) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}
