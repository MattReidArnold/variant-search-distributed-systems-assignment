package entities

import "time"

type Variant struct {
	Gene                   string
	NucleotideChange       string
	ProteinChange          string
	OtherMappings          string
	Alias                  string
	Transcripts            string
	Region                 string
	ReportedClassification string
	InferredClassification string
	Source                 string
	LastEvaluated          *time.Time
	LastUpdated            *time.Time
	URL                    string
	SubmitterComment       string
	Assembly               string
	Chr                    string
	GenomicStart           *int64
	GenomicStop            *int64
	Ref                    string
	Alt                    string
	ReportedRef            string
	ReportedAlt            string
}

func (v Variant) ToAvroMap() map[string]interface{} {
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
