package avro

import "github.com/mattreidarnold/variants/entities"

const VariantSchema = `{
	"type": "record",
	"name": "Variant",
	"fields": [
		{"name": "Gene", "type": "string"},
		{"name": "NucleotideChange", "type": "string"},
		{"name": "ProteinChange", "type": "string"},
		{"name": "OtherMappings", "type": "string"},
		{"name": "Alias", "type": "string"},
		{"name": "Transcripts", "type": "string"},
		{"name": "Region", "type": "string"},
		{"name": "ReportedClassification", "type": "string"},
		{"name": "InferredClassification", "type": "string"},
		{"name": "Source", "type": "string"},
		{"name": "LastEvaluated", "type": ["null",{ "type": "int", "logicalType":"date"}], "default":null},
		{"name": "LastUpdated", "type": ["null",{ "type": "int", "logicalType":"date"}], "default":null},
		{"name": "URL", "type": "string"},
		{"name": "SubmitterComment", "type": "string"},
		{"name": "Assembly", "type": "string"},
		{"name": "Chr", "type": "string"},
		{"name": "GenomicStart", "type": ["null", "int"], "default":null},
		{"name": "GenomicStop", "type": ["null", "int"], "default":null},
		{"name": "Ref", "type": "string"},
		{"name": "Alt", "type": "string"},
		{"name": "ReportedRef", "type": "string"},
		{"name": "ReportedAlt", "type": "string"}
	]
}`

func AvroMapFromVariant(v entities.Variant) map[string]interface{} {
	m := map[string]interface{}{
		"Gene":                   v.Gene,
		"NucleotideChange":       v.NucleotideChange,
		"ProteinChange":          v.ProteinChange,
		"OtherMappings":          v.OtherMappings,
		"Alias":                  v.Alias,
		"Transcripts":            v.Transcripts,
		"Region":                 v.Region,
		"ReportedClassification": v.ReportedClassification,
		"InferredClassification": v.InferredClassification,
		"Source":                 v.Source,
		"URL":                    v.URL,
		"SubmitterComment":       v.SubmitterComment,
		"Assembly":               v.Assembly,
		"Chr":                    v.Chr,
		"Ref":                    v.Ref,
		"Alt":                    v.Alt,
		"ReportedRef":            v.ReportedRef,
		"ReportedAlt":            v.ReportedAlt,
	}
	if v.LastEvaluated != nil {
		m["LastEvaluated"] = map[string]interface{}{
			"int.date": *v.LastEvaluated,
		}
	}
	if v.LastUpdated != nil {
		m["LastUpdated"] = map[string]interface{}{
			"int.date": *v.LastUpdated,
		}
	}
	if v.GenomicStart != nil {
		m["GenomicStart"] = map[string]interface{}{
			"int": *v.GenomicStart,
		}
	}
	if v.GenomicStop != nil {
		m["GenomicStop"] = map[string]interface{}{
			"int": *v.GenomicStop,
		}
	}
	return m
}
