package reporter

import (
	"encoding/json"

	"github.com/up9inc/oas-diff/differentiator"
)

type jsonReporter struct {
	changelog *differentiator.ChangelogOutput
}

func NewJSONReporter(changelog *differentiator.ChangelogOutput) Reporter {
	return &jsonReporter{
		changelog: changelog,
	}
}

func (j *jsonReporter) Build() ([]byte, error) {
	return json.MarshalIndent(j.changelog, "", "\t")
}
