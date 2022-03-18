package reporter

import (
	"bytes"
	"html/template"
	"sort"

	"github.com/up9inc/oas-diff/differentiator"
	"github.com/up9inc/oas-diff/model"
)

type htmlReporter struct {
	output *differentiator.ChangelogOutput
}

type pathChangelog struct {
	Key       string                   `json:"key"`
	Endpoint  string                   `json:"endpoint"`
	Operation string                   `json:"operation"`
	Changelog differentiator.Changelog `json:"changelog"`
}

type pathChangelogMap map[string][]pathChangelog

type ByType []pathChangelog

func (t ByType) Len() int           { return len(t) }
func (t ByType) Less(i, j int) bool { return t[i].Changelog.Type < t[j].Changelog.Type }
func (t ByType) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type ByOperation []*pathChangelog

func (o ByOperation) Len() int           { return len(o) }
func (o ByOperation) Less(i, j int) bool { return o[i].Operation < o[j].Operation }
func (o ByOperation) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

type templateData struct {
	Status            differentiator.ExecutionStatus
	Changelog         differentiator.ChangeMap
	PathChangelogList []pathChangelog
	PathChangelogMap  pathChangelogMap
}

func NewHTMLReporter(output *differentiator.ChangelogOutput) Reporter {
	return &htmlReporter{
		output: output,
	}
}

func (h *htmlReporter) Build() ([]byte, error) {
	data := templateData{
		Status:            h.output.ExecutionStatus,
		Changelog:         h.output.Changelog,
		PathChangelogList: h.buildPathChangelogList(),
		PathChangelogMap:  h.buildPathChangelogMap(),
	}

	var buf bytes.Buffer
	tmpl := template.Must(template.ParseFiles("reporter/template.html"))
	err := tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (h *htmlReporter) buildPathChangelogList() []pathChangelog {
	result := make([]pathChangelog, 0)

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

			result = append(result, pathChangelog{
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

func (h *htmlReporter) buildPathChangelogMap() pathChangelogMap {
	result := make(pathChangelogMap, 0)

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

			_, ok := result[endpoint]
			if !ok {
				result[endpoint] = make([]pathChangelog, 0)
			}
			result[endpoint] = append(result[endpoint], pathChangelog{
				Key:       k,
				Endpoint:  endpoint,
				Operation: op,
				Changelog: c,
			})
		}
	}

	// sort by type
	for _, v := range result {
		sort.Sort(ByType(v))
	}

	return result
}
