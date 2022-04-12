package reporter

import (
	"encoding/json"
	"sort"

	"github.com/up9inc/oas-diff/differentiator"
	"github.com/up9inc/oas-diff/model"
)

type endpoinChangelog struct {
	//Endpoint  string
	Operation string                   `json:"operation"`
	Changelog differentiator.Changelog `json:"changelog"`
}

type endpointData struct {
	Changelogs     []endpoinChangelog `json:"changelogs"`
	TotalChanges   int                `json:"total"`
	CreatedChanges int                `json:"created"`
	UpdatedChanges int                `json:"updated"`
	DeletedChanges int                `json:"deleted"`
}

type endpointsMap map[string]endpointData

type endpointReporter struct {
	output *differentiator.ChangelogOutput
}

func NewEndpointReporter(output *differentiator.ChangelogOutput) Reporter {
	return &endpointReporter{
		output: output,
	}
}

func (e *endpointReporter) Build() ([]byte, error) {
	return json.MarshalIndent(e.buildEndpointChangelogMap(), "", "\t")
}

func (e *endpointReporter) buildEndpointChangelogMap() endpointsMap {
	endpointsMap := make(endpointsMap, 0)

	for k, v := range e.output.Changelog {
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

			_, ok := endpointsMap[endpoint]
			if !ok {
				endpointsMap[endpoint] = endpointData{
					Changelogs:     make([]endpoinChangelog, 0),
					TotalChanges:   0,
					CreatedChanges: 0,
					UpdatedChanges: 0,
					DeletedChanges: 0,
				}
			}

			aux := endpointsMap[endpoint]
			aux.TotalChanges++
			switch c.Type {
			case "create":
				aux.CreatedChanges++
			case "update":
				aux.UpdatedChanges++
			case "delete":
				aux.DeletedChanges++
			}
			aux.Changelogs = append(endpointsMap[endpoint].Changelogs, endpoinChangelog{
				//Endpoint:  endpoint,
				Operation: op,
				Changelog: c,
			})
			endpointsMap[endpoint] = aux
		}
	}

	// sort by type
	for _, v := range endpointsMap {
		sort.Slice(v.Changelogs, func(i, j int) bool {
			return v.Changelogs[i].Changelog.Type < v.Changelogs[j].Changelog.Type
		})
	}

	return endpointsMap
}
