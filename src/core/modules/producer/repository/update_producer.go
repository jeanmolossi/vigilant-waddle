package repository

import (
	"context"
	"errors"

	"github.com/jeanmolossi/vigilant-waddle/src/domain/producer"
	"github.com/jeanmolossi/vigilant-waddle/src/pkg/drivers/database"
)

type updateProducer struct {
	db database.Database
}

// NewUpdateProducer will return a producer.UpdateProducerRepository implementation
func NewUpdateProducer(db database.Database) producer.UpdateProducerRepository {
	return &updateProducer{db}
}

// Run executes the update in producer
func (u *updateProducer) Run(ctx context.Context, usrID string, updater producer.ProducerUpdater) (producer.Producer, error) {
	// when has not user should fail.
	// that method only works if has usrID
	if usrID == "" {
		return nil, errors.New("fail on get usr with ID empty")
	}

	model := &ProducerModel{
		ID: usrID,
	}

	result := u.db.DB().First(model)
	if result.Error != nil {
		return nil, result.Error
	}

	// update the user that was recovered
	p, err := updater(model.AsDomain())
	if err != nil {
		return nil, err
	}

	// create a updated model to save
	// It should fill all fields to avoid
	// save data empty
	updatedModel := &ProducerModel{
		ID:        p.GetID(),
		Email:     p.GetEmail(),
		Password:  p.GetPassword(),
		Type:      p.GetScope().String(),
		CreatedAt: model.CreatedAt,
	}

	result = u.db.DB().Save(updatedModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return updatedModel.AsDomain(), nil
}
