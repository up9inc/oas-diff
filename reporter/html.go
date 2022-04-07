package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"sort"
	"strings"

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

var collapseHeaderIndex, collapseBodyIndex int

func (h *htmlReporter) Build() ([]byte, error) {
	funcMap := template.FuncMap{
		"CollapseHeaderId": func() string {
			collapseHeaderIndex++
			return fmt.Sprintf("collapseHeader_%d", collapseHeaderIndex)
		},
		"CollapseBodyId": func() string {
			collapseBodyIndex++
			return fmt.Sprintf("collapse_%d", collapseBodyIndex)
		},
		"StringLen": func(s string) int { return len(s) },
		"TotalPathsChanges": func(data pathChangelogMap) int {
			var count int
			for _, v := range data {
				count += len(v)
			}
			return count
		},
		"PathChanges": func(data []pathChangelog) int { return len(data) },
		"IsNotNil":    func(data interface{}) bool { return data != nil },
		"ToUpper":     strings.ToUpper,
		"ToLower":     strings.ToLower,
		"ToPrettyJSON": func(data interface{}) string {
			j, _ := json.MarshalIndent(data, "", "\t")
			return string(j)
		},
		"FormatPath": func(path []string) string { return strings.Join(path, " ") },
		"GetTypeInfo": func(t string, s differentiator.ExecutionStatus) string {
			switch t {
			case "create":
				return fmt.Sprintf("CREATED from %s", s.SecondFilePath)
			// TODO: Updated from source file info
			case "update":
				return "UPDATED"
			case "delete":
				return fmt.Sprintf("DELETED from %s", s.BaseFilePath)
			}

			return ""
		},
		"GetTypeColor": func(t string) string {
			switch t {
			case "create":
				return "success"
			case "update":
				return "warning"
			case "delete":
				return "danger"
			}

			return ""
		},
		"GetFromTypeColor": func(t string) string {
			switch t {
			case "create":
				return "success"
			case "update":
				return "danger"
			case "delete":
				return "danger"
			}

			return "info"
		},
		"GetToTypeColor": func(t string) string {
			switch t {
			case "create":
				return "success"
			case "update":
				return "success"
			case "delete":
				return "danger"
			}

			return "info"
		},
	}

	data := templateData{
		Status:            h.output.ExecutionStatus,
		Changelog:         h.output.Changelog,
		PathChangelogList: h.buildPathChangelogList(),
		PathChangelogMap:  h.buildPathChangelogMap(),
	}

	var buf bytes.Buffer
	tmpl, err := template.New("template.html").Funcs(funcMap).ParseFiles("reporter/template.html")
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&buf, data)
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
			var endpoint string

			if len(c.Path) > 0 {
				endpoint = c.Path[0]
			}
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
			var endpoint string

			if len(c.Path) > 0 {
				endpoint = c.Path[0]
			}
			if len(c.Path) > 1 {
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
