// Package producer is a domain package wich contains the producer features and behavior.
package producer

import baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"

// Producer has the baseUser interface implemented
type Producer interface {
	baseuser.BaseUser
}

// Usecases is the producer usecases

// RegisterProducer is the usecase to register a new producer
type RegisterProducer func(Producer) (Producer, error)

// GetMe is the usecase to get the current producer
// It will use the session token
type GetMe func(GetMeOptions) (Producer, error)

// GetMeOptions is the options for the GetMe usecase
type GetMeOptions struct {
	ProducerID string   `json:"producer_id" example:"1" format:"uuid" validate:"required,uuid"`
	Fields     []string `query:"fields" example:"id,email"`
}

// GetErrorMap implements ModuleErrorMap to GetMeOptions
func (g *GetMeOptions) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"producerid": {
			"required": ErrMissingProducerID,
			"uuid":     ErrProducerIDInvalid},
	}
}
