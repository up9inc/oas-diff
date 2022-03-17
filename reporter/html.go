package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/up9inc/oas-diff/differentiator"
	"github.com/up9inc/oas-diff/model"
)

type htmlReporter struct {
	output *differentiator.ChangelogOutput
}

type changelogByPath struct {
	Key       string                    `json:"key"`
	Path      string                    `json:"path"`
	Operation string                    `json:"operation"`
	Changelog *differentiator.Changelog `json:"changelog"`
}

type changelogByType struct {
	Created []*changelogByPath `json:"created"`
	Updated []*changelogByPath `json:"updated"`
	Deleted []*changelogByPath `json:"deleted"`
}

type templateData struct {
	Status          differentiator.ExecutionStatus
	Changelog       string
	ChangelogByType string
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
	changelogByTypeJSON, err := json.Marshal(h.groupByType())
	if err != nil {
		return nil, err
	}

	data := templateData{
		Status:          h.output.ExecutionStatus,
		Changelog:       string(changelogJSON),
		ChangelogByType: string(changelogByTypeJSON),
	}

	var buf bytes.Buffer
	tmpl := template.Must(template.ParseFiles("reporter/template.html"))
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (h *htmlReporter) groupByType() *changelogByType {
	pathGroup := h.groupByPath()
	created := make([]*changelogByPath, 0)
	updated := make([]*changelogByPath, 0)
	deleted := make([]*changelogByPath, 0)

	for _, p := range pathGroup {
		switch p.Changelog.Type {
		case "create":
			created = append(created, p)
		case "update":
			updated = append(updated, p)
		case "delete":
			deleted = append(deleted, p)
		default:
			panic(fmt.Sprintf("invalid changelog type: %s", p.Changelog.Type))
		}
	}

	return &changelogByType{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}
}

func (h *htmlReporter) groupByPath() []*changelogByPath {
	result := make([]*changelogByPath, 0)

	for k, v := range h.output.Changelog {
		for _, c := range v {
			// TODO: Default value for operation when we don't have the operation method in path
			var path, op string
			if k == model.OAS_PATHS_KEY || k == model.OAS_WEBHOOKS_KEY {
				if len(c.Path) > 1 && c.Path[1] != "parameters" {
					path = c.Path[0]
					op = c.Path[1]
				}
			}

			result = append(result, &changelogByPath{
				Key:       k,
				Path:      path,
				Operation: op,
				Changelog: c,
			})
		}
	}

	return result
}
