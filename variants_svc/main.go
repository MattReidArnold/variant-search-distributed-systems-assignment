package main

import (
	"fmt"
	"log"
	"os"

	kafkaavro "github.com/dangkaka/go-kafka-avro"

	usecases "github.com/mattreidarnold/variants/application/use_cases"
	"github.com/mattreidarnold/variants/entities"
	"github.com/mattreidarnold/variants/external_interfaces/kafka"
	"github.com/mattreidarnold/variants/external_interfaces/tsv"
)

func main() {

	// var kafkaServers = []string{"localhost:9092"}
	// var schemaRegistryServers = []string{"http://localhost:8081"}
	// var topic = "variants"
	kafkaServers := []string{os.Getenv("KAFKA_BROKER")}
	schemaRegistryServers := []string{os.Getenv("SCHEMA_REGISTRY")}
	topic := os.Getenv("VARIANTS_CDC_TOPIC")
	dataFile := os.Getenv("VARIANTS_DATA_FILE")

	producer, err := kafkaavro.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		log.Fatalf("Could not create avro producer: %s", err)
	}

	variantProducer := kafka.NewVariantUpdateEventProducer(producer, topic)

	updateVariant := usecases.NewUpdateVariantCommand(variantProducer)

	err = tsv.Read(dataFile, rowHandler(updateVariant), func(rowsProcessed int, errs []error) error {
		fmt.Printf("Processed %v rows with %v errors\n", rowsProcessed, len(errs))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func rowHandler(updateVariant usecases.UpdateVariantCommand) func(r []string) error {
	return func(row []string) error {
		v, err := rowToVariant(row)
		if err != nil {
			return err
		}
		return updateVariant.Execute(v)
	}
}

func rowToVariant(row []string) (entities.Variant, error) {
	v := entities.Variant{
		Gene:                   row[0],
		NucleotideChange:       row[1],
		ProteinChange:          row[2],
		OtherMappings:          row[3],
		Alias:                  row[4],
		Transcripts:            row[5],
		Region:                 row[6],
		ReportedClassification: row[7],
		InferredClassification: row[8],
		Source:                 row[9],
		URL:                    row[12],
		SubmitterComment:       row[13],
		Assembly:               row[14],
		Chr:                    row[15],
		Ref:                    row[18],
		Alt:                    row[19],
		ReportedRef:            row[20],
		ReportedAlt:            row[21],
	}

	lastEvaluated, err := tsv.ParseTime(row[10])
	if err != nil {
		return entities.Variant{}, err
	}
	v.LastEvaluated = lastEvaluated

	lastUpdated, err := tsv.ParseTime(row[11])
	if err != nil {
		return entities.Variant{}, err
	}
	v.LastUpdated = lastUpdated

	gStart, err := tsv.ParseInt(row[16])
	if err != nil {
		return entities.Variant{}, err
	}
	v.GenomicStart = gStart

	gStop, err := tsv.ParseInt(row[17])
	if err != nil {
		return entities.Variant{}, err
	}
	v.GenomicStart = gStop

	return v, nil
}
