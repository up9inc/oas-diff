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

type pathData struct {
	Paths         []pathChangelog `json:"paths"`
	TotalChanges  int             `json:"totalChanges"`
	CreateChanges int             `json:"createChanges"`
	UpdateChanges int             `json:"updateChanges"`
	DeleteChanges int             `json:"deleteChanges"`
}

type pathChangelogMap map[string]pathData

type ByType []pathChangelog

func (t ByType) Len() int           { return len(t) }
func (t ByType) Less(i, j int) bool { return t[i].Changelog.Type < t[j].Changelog.Type }
func (t ByType) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type templateData struct {
	Status               differentiator.ExecutionStatus
	NonPathChangelogList []differentiator.Changelog
	PathChangelogMap     pathChangelogMap
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
				count += v.TotalChanges
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
		Status:               h.output.ExecutionStatus,
		NonPathChangelogList: h.buildNonPathChangelogList(),
		PathChangelogMap:     h.buildPathChangelogMap(),
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
							//c.Path = append(c.Path, op)
						}
					}
				}
			}

			_, ok := result[endpoint]
			if !ok {
				result[endpoint] = pathData{
					TotalChanges:  0,
					CreateChanges: 0,
					UpdateChanges: 0,
					DeleteChanges: 0,
					Paths:         make([]pathChangelog, 0),
				}
			}

			aux := result[endpoint]
			aux.TotalChanges++
			switch c.Type {
			case "create":
				aux.CreateChanges++
			case "update":
				aux.UpdateChanges++
			case "delete":
				aux.DeleteChanges++
			}
			aux.Paths = append(result[endpoint].Paths, pathChangelog{
				Key:       k,
				Endpoint:  endpoint,
				Operation: op,
				Changelog: c,
			})
			result[endpoint] = aux
		}
	}

	// sort by type
	for _, v := range result {
		sort.Sort(ByType(v.Paths))
	}

	return result
}
