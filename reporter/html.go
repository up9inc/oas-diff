package reporter

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/up9inc/oas-diff/differentiator"
)

type htmlReporter struct {
	output *differentiator.ChangelogOutput
}

type templateData struct {
	Status    differentiator.ExecutionStatus
	Changelog string
}

func NewHTMLReporter(output *differentiator.ChangelogOutput) Reporter {
	return &htmlReporter{
		output: output,
	}
}

func (h *htmlReporter) Build() ([]byte, error) {
	// The issue with passing struct is that pointers won't be dereferenced
	// Let's pass a JSON to the html template instead of a struct
	changelogJSON, err := json.Marshal(h.output.Changelog)
	if err != nil {
		return nil, err
	}
	data := templateData{
		Status:    h.output.ExecutionStatus,
		Changelog: string(changelogJSON),
	}

	var buf bytes.Buffer
	tmpl := template.Must(template.ParseFiles("reporter/template.html"))
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
