package reporter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/util"
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

type SummaryData struct {
	Endpoints      map[string][]string
	RequestHeaders map[string]map[string][]string
	Parameters     map[string]map[string][]string
}

func (s *SummaryData) AddEndpoint(typeKey string, value string) {
	_, ok := s.Endpoints[typeKey]
	if !ok {
		s.Endpoints[typeKey] = make([]string, 0)
	}
	s.Endpoints[typeKey] = util.SliceElementAddUnique(s.Endpoints[typeKey], value)
}

func (s *SummaryData) AddRequestHeader(typeKey string, endpointKey, value string) {
	_, ok := s.RequestHeaders[typeKey]
	if !ok {
		s.RequestHeaders[typeKey] = make(map[string][]string, 0)
	}
	_, ok = s.RequestHeaders[typeKey][endpointKey]
	if !ok {
		s.RequestHeaders[typeKey][endpointKey] = make([]string, 0)
	}

	s.RequestHeaders[typeKey][endpointKey] = util.SliceElementAddUnique(s.RequestHeaders[typeKey][endpointKey], value)
}

func (s *SummaryData) AddParameter(typeKey string, endpointKey, value string) {
	_, ok := s.Parameters[typeKey]
	if !ok {
		s.Parameters[typeKey] = make(map[string][]string, 0)
	}
	_, ok = s.Parameters[typeKey][endpointKey]
	if !ok {
		s.Parameters[typeKey][endpointKey] = make([]string, 0)
	}

	s.Parameters[typeKey][endpointKey] = util.SliceElementAddUnique(s.Parameters[typeKey][endpointKey], value)
}

func (s *summaryReporter) buildEndpointChangelogMap() (SummaryData, error) {
	params := model.Parameters{}
	summaryData := SummaryData{
		Endpoints:      make(map[string][]string, 0),
		RequestHeaders: make(map[string]map[string][]string, 0),
		Parameters:     make(map[string]map[string][]string, 0),
	}

	for k, v := range s.output.Changelog {
		for _, c := range v {
			// ignore others non-paths keys
			if k != model.OAS_PATHS_KEY && k != model.OAS_WEBHOOKS_KEY {
				continue
			}

			var operation *model.Operation
			var op string
			var endpoint string
			var typeKey string
			var endpointKey string
			pathItem := &model.PathItem{}

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
				typeKey = "New"
				// file2
				sourceFileRef = s.jsonFile2
			case "delete":
				typeKey = "Removed"
				// file1
				sourceFileRef = s.jsonFile
			case "update":
				typeKey = "Updated"
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
				return summaryData, err
			}

			if len(c.Path) > 1 {
				op = c.Path[1]
				endpointKey = fmt.Sprintf("%s %s", strings.ToUpper(op), endpoint)

				// updated endpoints only
				if c.Type == "update" {
					summaryData.AddEndpoint(typeKey, endpointKey)
				}

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
								summaryData.AddRequestHeader(typeKey, endpointKey, pv.Name)
								//summaryData.AddRequestHeader(typeKey, endpointKey, c.Identifier[params.GetIdentifierName()])
							} else {
								summaryData.AddParameter(typeKey, endpointKey, pv.Name)
								//summaryData.AddParameter(typeKey, endpointKey, c.Identifier[params.GetIdentifierName()])
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
							endpointKey = fmt.Sprintf("%s %s", strings.ToUpper(op), endpoint)

							// created/deleted endpoints only
							summaryData.AddEndpoint(typeKey, endpointKey)
						}
					}
				}
			}
		}
	}

	return summaryData, nil
}
