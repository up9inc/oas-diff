package reporter

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
)

type endpoinChangelog struct {
	Type       string                   `json:"type"`
	Endpoint   string                   `json:"endpoint"`
	Operation  string                   `json:"operation"`
	Headers    []string                 `json:"headers"`
	Parameters []string                 `json:"parameters"`
	Changelog  differentiator.Changelog `json:"changelog"`
}

type endpointData struct {
	Changelogs     []endpoinChangelog `json:"changelogs"`
	TotalChanges   int                `json:"total"`
	CreatedChanges int                `json:"created"`
	UpdatedChanges int                `json:"updated"`
	DeletedChanges int                `json:"deleted"`
}

type endpointsMap map[string]endpointData

type summaryReporter struct {
	output    *differentiator.ChangelogOutput
	jsonFile  file.JsonFile
	jsonFile2 file.JsonFile
}

func NewSummaryReporter(jsonFile file.JsonFile, jsonFile2 file.JsonFile, output *differentiator.ChangelogOutput) Reporter {
	return &summaryReporter{
		output:    output,
		jsonFile:  jsonFile,
		jsonFile2: jsonFile2,
	}
}

func (s *summaryReporter) Build() ([]byte, error) {
	data, err := s.buildEndpointChangelogMap()
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(data, "", "\t")
}

func (s *summaryReporter) buildEndpointChangelogMap() (endpointsMap, error) {
	endpointsMap := make(endpointsMap, 0)
	params := model.Parameters{}

	for k, v := range s.output.Changelog {
		for _, c := range v {
			// ignore others non-paths keys
			if k != model.OAS_PATHS_KEY && k != model.OAS_WEBHOOKS_KEY {
				continue
			}

			pathItem := &model.PathItem{}
			var operation *model.Operation
			var op string
			var endpoint string
			var headers []string
			var parameters []string

			if len(c.Path) > 0 {
				endpoint = c.Path[0]
			}

			if len(endpoint) == 0 {
				panic("endpoint should not be nil")
			}

			var sourceFileRef file.JsonFile
			var endpointNode *[]byte

			switch c.Type {
			case "create":
				// file2
				sourceFileRef = s.jsonFile2
			case "delete":
				// file1
				sourceFileRef = s.jsonFile
			case "update":
				// both
				// try file1 first
				sourceFileRef = s.jsonFile
				endpointNode = sourceFileRef.GetNodeData(fmt.Sprintf("%s.%s", model.OAS_PATHS_KEY, endpoint))
				if endpointNode == nil {
					sourceFileRef = s.jsonFile2
				}
			}

			if endpointNode == nil {
				endpointNode = sourceFileRef.GetNodeData(fmt.Sprintf("%s.%s", model.OAS_PATHS_KEY, endpoint))
				if endpointNode == nil {
					panic(fmt.Errorf(`failed to find endpoint "%s" node for "%s" operation in file "%s"`, endpoint, c.Type, sourceFileRef.GetPath()))
				}
			}
			err := pathItem.ParseFromNode(endpointNode)
			if err != nil {
				return nil, err
			}

			if len(c.Path) > 1 {
				op = c.Path[1]

				// TODO: How to distinguish Parameters type: "query", "header", "path" or "cookie"
				// TODO: Response Headers Map

				// endpoint.parameters || endpoint.operation.parameters
				if op == params.GetName() || (len(c.Path) > 2 && c.Path[2] == params.GetName()) {
					paramsRef := pathItem.Parameters

					// endpoint.operation.parameters
					if op != params.GetName() {
						switch op {
						case "connect":
							operation = pathItem.Connect
						case "delete":
							operation = pathItem.Delete
						case "get":
							operation = pathItem.Get
						case "head":
							operation = pathItem.Head
						case "options":
							operation = pathItem.Options
						case "patch":
							operation = pathItem.Patch
						case "post":
							operation = pathItem.Post
						case "put":
							operation = pathItem.Put
						case "trace":
							operation = pathItem.Trace
						}

						if operation == nil {
							panic("operation should not be nil")
						}

						paramsRef = operation.Parameters
					}

					for _, pv := range paramsRef {
						if pv.Name == c.Identifier[params.GetIdentifierName()] {
							paramType := pv.In

							// endpoint.operation.parameters.identifier.in || endpoint.parameters.identifier.in
							if (len(c.Path) > 4 && c.Path[4] == "in") || (op == params.GetName() && len(c.Path) > 3 && c.Path[3] == "in") {
								paramType = c.To.(string)
							}

							if paramType == "header" {
								headers = append(headers, pv.Name)
								//headers = append(headers, c.Identifier[params.GetIdentifierName()])
							} else {
								parameters = append(parameters, pv.Name)
								//parameters = append(parameters, c.Identifier[params.GetIdentifierName()])
							}
							break
						}
					}

				}

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
				Type:       c.Type,
				Endpoint:   endpoint,
				Operation:  op,
				Headers:    headers,
				Parameters: parameters,
				Changelog:  c,
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

	return endpointsMap, nil
}
