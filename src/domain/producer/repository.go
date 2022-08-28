package producer

import "context"

// RegisterProducerRepository is the repository implementation for Register Producer.
type RegisterProducerRepository interface {
	// Run executes the Register Producer query.
	//
	// Example:
	//
	//      newProducer, err := repo.Run(context.Backround(), producer)
	//
	Run(ctx context.Context, p Producer) (Producer, error)
}

// ProducerUpdater is a interface who defines the UpdateProducerRepository method
type ProducerUpdater func(Producer) (Producer, error)

// UpdateProducerRepository is repository implementation for Update Producer.
type UpdateProducerRepository interface {
	// Run executes the update in producer
	Run(ctx context.Context, usrID string, u ProducerUpdater) (Producer, error)
}
