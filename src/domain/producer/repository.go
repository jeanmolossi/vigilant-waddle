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
