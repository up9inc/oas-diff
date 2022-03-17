package reporter

import (
	"encoding/json"

	"github.com/up9inc/oas-diff/differentiator"
)

type jsonReporter struct {
	output *differentiator.ChangelogOutput
}

func NewJSONReporter(output *differentiator.ChangelogOutput) Reporter {
	return &jsonReporter{
		output: output,
	}
}

func (j *jsonReporter) Build() ([]byte, error) {
	return json.MarshalIndent(j.output, "", "\t")
}
