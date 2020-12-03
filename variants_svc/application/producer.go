package application

import "github.com/mattreidarnold/variants/entities"

type VariantUpdateEventProducer interface {
	Produce(v entities.Variant) error
}
