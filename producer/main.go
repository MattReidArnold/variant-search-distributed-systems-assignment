package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/dangkaka/go-kafka-avro"
)

type Variant struct {
	Gene             string
	NucleotideChange string
	// OtherMappings          string
	// Transcripts            string
	// Region                 string
	// ReportedClassification string
	// Source                 string
	// LastEvaluated          time.Time
	// LastUpdated            time.Time
	// URL                    string
	// SubmitterComment       string
	// Assembly               string
	// Chr                    string
	// GenomicStart           int32
	// GenomicStop            int32
	// Ref                    string
	// Alt                    string
	// ReportedRef            string
	// ReportedAlt            string
}

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "variants"
var schema = `{
	"type": "record",
	"name": "Variant",
	"fields": [
		{"name": "Gene", "type": "string"},
		{"name": "NucleotideChange", "type": "string"}
	]
}`

func main() {
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		log.Fatalf("Could not create avro producer: %s", err)
	}

	err = readCsv("data/variants-new.tsv", '\t', func(row []string) error {
		v, err := rowToVariant(row)
		if err != nil {
			return err
		}

		value, err := json.Marshal(v)
		if err != nil {
			return err
		}

		key := time.Now().String()
		err = producer.Add(topic, schema, []byte(key), []byte(value))

		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

func rowToVariant(row []string) (Variant, error) {
	v := Variant{
		Gene:             row[0],
		NucleotideChange: row[1],
	}
	return v, nil
}

type rowCallback func(row []string) error

func readCsv(filename string, delimeter rune, cb rowCallback) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = delimeter

	headerRow, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Printf("Processing file with headers: %s", headerRow)

	for i := 0; i < 10; i++ {
		row, err := r.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = cb(row)
		if err != nil {
			return err
		}
	}
	return nil
}
