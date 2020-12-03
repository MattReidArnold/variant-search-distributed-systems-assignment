package kafka

import (
	"crypto/md5"
	"fmt"

	kafkaavro "github.com/dangkaka/go-kafka-avro"
	"github.com/linkedin/goavro"

	"github.com/mattreidarnold/variants/application"
	"github.com/mattreidarnold/variants/entities"
	"github.com/mattreidarnold/variants/external_interfaces/avro"
)

type variantUpdateEventProducer struct {
	producer *kafkaavro.AvroProducer
	topic    string
}

func NewVariantUpdateEventProducer(producer *kafkaavro.AvroProducer, topic string) application.VariantUpdateEventProducer {
	return &variantUpdateEventProducer{
		producer: producer,
		topic:    topic,
	}
}

func (p *variantUpdateEventProducer) Produce(v entities.Variant) error {
	codec, err := goavro.NewCodec(avro.VariantSchema)
	if err != nil {
		return err
	}
	value, err := codec.TextualFromNative(nil, avro.AvroMapFromVariant(v))
	if err != nil {
		return err
	}
	key := generateVariantID(v)

	err = p.producer.Add(p.topic, avro.VariantSchema, []byte(key), value)
	return nil
}

func generateVariantID(v entities.Variant) string {
	data := []byte(v.Gene + v.NucleotideChange + v.OtherMappings + v.Transcripts + v.Region)
	return fmt.Sprintf("%x", md5.Sum(data))
}
