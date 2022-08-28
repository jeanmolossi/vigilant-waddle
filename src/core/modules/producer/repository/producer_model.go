package repository

import (
	"time"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"gorm.io/gorm"
)

// producersTable is a pkg var to retrieve Producers table name
// var producersTable = ProducerModel{}.TableName()

// ProducerModel represents how a producer is stored in the database.
type ProducerModel struct {
	ID        string    `gorm:"primary_key;column:usr_id;type:uuid;default:uuid_generate_v4();index:idx_usr_id"`
	Email     string    `gorm:"column:usr_email;unique;type:varchar(255);"`
	Password  string    `gorm:"column:usr_password;type:varchar(255)"`
	Type      string    `gorm:"column:usr_scope;type:varchar(64);default:student"`
	CreatedAt time.Time `gorm:"column:usr_created_at;type:timestamp with time zone;default:now()"`
	UpdatedAt time.Time `gorm:"column:usr_updated_at;type:timestamp with time zone;default:now()"`
}

// TableName overrides the table name used by ProducerModel to `users`
//
// Read more about GORM conventions:
//
// https://gorm.io/docs/conventions.html#TableName
func (ProducerModel) TableName() string {
	return "users"
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
//
// Read more about GORM hooks:
//
// https://gorm.io/docs/hooks.html#Creating-an-object
func (s *ProducerModel) BeforeCreate(*gorm.DB) error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
//
// Read more about GORM hooks:
//
// https://gorm.io/docs/hooks.html#Updating-an-object
func (s *ProducerModel) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *ProducerModel) AsDomain() producer.Producer {
	return producer.NewProducer(
		producer.WithID(s.ID),
		producer.WithEmail(s.Email),
		producer.WithPassword(s.Password),
		producer.WithScope(s.Type),
	)
}
