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
	Endpoint  string
	Operation string
	Changelog differentiator.Changelog
}

type pathsData struct {
	Key           string
	Paths         []pathChangelog
	TotalChanges  int
	CreateChanges int
	UpdateChanges int
	DeleteChanges int
}

type pathsMap map[string]pathsData

type byType []pathChangelog

func (t byType) Len() int           { return len(t) }
func (t byType) Less(i, j int) bool { return t[i].Changelog.Type < t[j].Changelog.Type }
func (t byType) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type pathKeyValue struct {
	Key   string
	Value pathsData
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
		"TotalPathsChanges": func(data []pathKeyValue) int {
			var count int
			for _, v := range data {
				count += v.Value.TotalChanges
			}
			return count
		},
		"IsNotNil": func(data interface{}) bool { return data != nil },
		"ToUpper":  strings.ToUpper,
		"ToLower":  strings.ToLower,
		"ToPrettyJSON": func(data interface{}) string {
			j, _ := json.MarshalIndent(data, "", "\t")
			return string(j)
		},
		"FormatPath":  func(path []string) string { return strings.Join(path, " ") },
		"PathPadding": func(index int) string { return fmt.Sprintf("padding-left: %.1fem", float32(index)*0.4) },
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

	data := struct {
		Status               differentiator.ExecutionStatus
		NonPathChangelogList []differentiator.Changelog
		PathChangelogList    []pathKeyValue
	}{
		Status:               h.output.ExecutionStatus,
		NonPathChangelogList: h.buildNonPathChangelogList(),
		PathChangelogList:    h.buildPathChangelogMap(),
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

func (h *htmlReporter) buildNonPathChangelogList() []differentiator.Changelog {
	result := make([]differentiator.Changelog, 0)

	for k, v := range h.output.Changelog {
		for _, c := range v {

			// ignore paths and webhooks
			if k == model.OAS_PATHS_KEY || k == model.OAS_WEBHOOKS_KEY {
				continue
			}

			result = append(result, c)
		}
	}

	// sort by type
	sort.Sort(differentiator.ByType(result))

	return result
}

func (h *htmlReporter) buildPathChangelogMap() []pathKeyValue {
	pathsMap := make(pathsMap, 0)

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
			} else {
				// the endpoint was created/deleted, we only have one operation
				if c.Type != "update" {
					data := c.To
					if c.Type == "delete" {
						data = c.From
					}
					pathItem, ok := data.(model.PathItem)
					if ok {
						operations := pathItem.GetOperationsName()
						if len(operations) == 1 {
							op = operations[0]
							c.Path = append(c.Path, op)
						}
					}
				}
			}

			_, ok := pathsMap[endpoint]
			if !ok {
				pathsMap[endpoint] = pathsData{
					Key:           k,
					Paths:         make([]pathChangelog, 0),
					TotalChanges:  0,
					CreateChanges: 0,
					UpdateChanges: 0,
					DeleteChanges: 0,
				}
			}

			aux := pathsMap[endpoint]
			aux.TotalChanges++
			switch c.Type {
			case "create":
				aux.CreateChanges++
			case "update":
				aux.UpdateChanges++
			case "delete":
				aux.DeleteChanges++
			}
			aux.Paths = append(pathsMap[endpoint].Paths, pathChangelog{
				Endpoint:  endpoint,
				Operation: op,
				Changelog: c,
			})
			pathsMap[endpoint] = aux
		}
	}

	var ss []pathKeyValue

	// sort by type
	for k, v := range pathsMap {
		sort.Sort(byType(v.Paths))
		ss = append(ss, pathKeyValue{k, v})
	}

	// sorty by TotalChanges
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value.TotalChanges > ss[j].Value.TotalChanges
	})

	return ss
}
