package reporter

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"sort"
	"text/template"

	"github.com/up9inc/oas-diff/differentiator"
	"github.com/up9inc/oas-diff/model"
)

//go:embed template.html
var templateFS embed.FS

const templateName = "template.html"

type htmlReporter struct {
	output *differentiator.ChangelogOutput
}

type pathChangelog struct {
	Endpoint  string                   `json:"endpoint"`
	Operation string                   `json:"operation"`
	Changelog differentiator.Changelog `json:"changelog"`
}

type pathsData struct {
	Key            string          `json:"key"`
	Paths          []pathChangelog `json:"path"`
	TotalChanges   int             `json:"totalChanges"`
	CreatedChanges int             `json:"createdChanges"`
	UpdatedChanges int             `json:"updatedChanges"`
	DeletedChanges int             `json:"deletedChanges"`
}

type pathsMap map[string]pathsData

type byType []pathChangelog

func (t byType) Len() int           { return len(t) }
func (t byType) Less(i, j int) bool { return t[i].Changelog.Type < t[j].Changelog.Type }
func (t byType) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type pathKeyValue struct {
	Key   string    `json:"key"`
	Value pathsData `json:"value"`
}

func NewHTMLReporter(output *differentiator.ChangelogOutput) Reporter {
	return &htmlReporter{
		output: output,
	}
}

var collapseHeaderIndex, collapseBodyIndex int

func (h *htmlReporter) Build() ([]byte, error) {
	buildPathChangelogJson, err := json.Marshal(h.buildPathChangelogMap())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	buildStatusJson, err := json.Marshal(h.output.ExecutionStatus)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data := struct {
		Status            string
		PathChangelogList string
	}{
		Status:            string(buildStatusJson),
		PathChangelogList: string(buildPathChangelogJson),
	}

	var buf bytes.Buffer
	tmpl, err := template.New(templateName).ParseFS(templateFS, templateName)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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
					Key:            k,
					Paths:          make([]pathChangelog, 0),
					TotalChanges:   0,
					CreatedChanges: 0,
					UpdatedChanges: 0,
					DeletedChanges: 0,
				}
			}

			aux := pathsMap[endpoint]
			aux.TotalChanges++
			switch c.Type {
			case "create":
				aux.CreatedChanges++
			case "update":
				aux.UpdatedChanges++
			case "delete":
				aux.DeletedChanges++
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
