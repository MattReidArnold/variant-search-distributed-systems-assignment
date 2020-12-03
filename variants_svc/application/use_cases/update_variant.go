package usecases

import (
	"github.com/mattreidarnold/variants/application"
	"github.com/mattreidarnold/variants/entities"
)

type UpdateVariantCommand interface {
	Execute(v entities.Variant) error
}

type updateVariantCommand struct {
	producer application.VariantUpdateEventProducer
}

func NewUpdateVariantCommand(producer application.VariantUpdateEventProducer) UpdateVariantCommand {
	return &updateVariantCommand{
		producer: producer,
	}
}

func (c *updateVariantCommand) Execute(v entities.Variant) error {
	return c.producer.Produce(v)
}
