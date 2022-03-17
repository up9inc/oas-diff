package reporter

import (
	"bytes"
	"encoding/json"
	"html/template"
	"sort"

	"github.com/up9inc/oas-diff/differentiator"
	"github.com/up9inc/oas-diff/model"
)

type htmlReporter struct {
	output *differentiator.ChangelogOutput
}

type pathChangelog struct {
	Key       string                    `json:"key"`
	Endpoint  string                    `json:"endpoint"`
	Operation string                    `json:"operation"`
	Changelog *differentiator.Changelog `json:"changelog"`
}

type ByType []*pathChangelog

func (t ByType) Len() int           { return len(t) }
func (t ByType) Less(i, j int) bool { return t[i].Changelog.Type < t[j].Changelog.Type }
func (t ByType) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type templateData struct {
	Status        differentiator.ExecutionStatus
	Changelog     string
	PathChangelog string
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

	pathChangelogJSON, err := json.Marshal(h.buildPathChangelog())
	if err != nil {
		return nil, err
	}

	data := templateData{
		Status:        h.output.ExecutionStatus,
		Changelog:     string(changelogJSON),
		PathChangelog: string(pathChangelogJSON),
	}

	var buf bytes.Buffer
	tmpl := template.Must(template.ParseFiles("reporter/template.html"))
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (h *htmlReporter) buildPathChangelog() []*pathChangelog {
	result := make([]*pathChangelog, 0)

	for k, v := range h.output.Changelog {
		for _, c := range v {
			// ignore others non-paths keys
			if k != model.OAS_PATHS_KEY && k != model.OAS_WEBHOOKS_KEY {
				continue
			}

			var op string
			endpoint := c.Path[0]
			if len(c.Path) > 1 && c.Path[1] != "parameters" {
				op = c.Path[1]
			}

			result = append(result, &pathChangelog{
				Key:       k,
				Endpoint:  endpoint,
				Operation: op,
				Changelog: c,
			})
		}
	}

	// sort by type
	sort.Sort(ByType(result))

	return result
}
