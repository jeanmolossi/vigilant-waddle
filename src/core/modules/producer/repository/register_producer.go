package repository

import (
	"context"

	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type registerProducer struct {
	db database.Database
}

func NewRegisterProducer(db database.Database) producer.RegisterProducerRepository {
	return &registerProducer{db}
}

func (r *registerProducer) Run(ctx context.Context, p producer.Producer) (producer.Producer, error) {
	_producer := &ProducerModel{
		ID:       p.GetID(),
		Email:    p.GetEmail(),
		Password: p.GetPassword(),
		Type:     baseuser.PRODUCER.String(),
	}

	result := r.db.DB().Create(_producer)
	if result.Error != nil {
		return nil, result.Error
	}

	newProducer := producer.NewProducer(
		producer.WithID(_producer.ID),
		producer.WithEmail(_producer.Email),
		producer.WithScope(_producer.Type),
	)

	return newProducer, nil
}
