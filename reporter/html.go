package reporter

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/up9inc/oas-diff/differentiator"
)

type htmlReporter struct {
	changelog *differentiator.ChangelogOutput
}

func NewHTMLReporter(changelog *differentiator.ChangelogOutput) Reporter {
	return &htmlReporter{
		changelog: changelog,
	}
}

func (h *htmlReporter) Build() ([]byte, error) {
	// The issue with passing struct is that pointers won't be dereferenced
	// Let's pass a JSON to the html template instead of a struct
	data, err := json.Marshal(h.changelog)
	if err != nil {
		return nil, err
	}

	tmpl := template.Must(template.ParseFiles("reporter/template.html"))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, string(data))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
