package differentiator

import (
	"fmt"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*pathsDiff)(nil)

type pathsDiff struct {
	*internalDiff
	data  model.Paths
	data2 model.Paths
}

func NewPathsDiff() *pathsDiff {
	return &pathsDiff{
		internalDiff: NewInternalDiff(model.OAS_PATHS_KEY),
		data:         model.Paths{},
		data2:        model.Paths{},
	}
}

func (p *pathsDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) error {
	var err error

	// schema
	err = p.schema.Build(validator)
	if err != nil {
		return err
	}

	// paths1
	p.filePath = jsonFile.GetPath()
	err = p.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// paths2
	p.filePath2 = jsonFile2.GetPath()
	err = p.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	/* 	// TODO: Diff segments of the paths instead the entire map
	   	for k := range p.data {
	   		// Parameters
	   		err = p.handleArray(p.data[k].Parameters, p.data2[k].Parameters)
	   		if err != nil {
	   			return err
	   		}
	   		p.data[k].Parameters = nil
	   		p.data2[k].Parameters = nil

	   		// Servers
	   		err = p.handleArray(p.data[k].Servers, p.data2[k].Servers)
	   		if err != nil {
	   			return err
	   		}
	   		p.data[k].Servers = nil
	   		p.data2[k].Servers = nil

	   		// Operations

	   		// Connect
	   		if p.data[k].Connect != nil && p.data2[k].Connect != nil {
	   			// Connect Parameters
	   			err = p.handleArray(p.data[k].Connect.Parameters, p.data2[k].Connect.Parameters)
	   			if err != nil {
	   				return err
	   			}
	   			p.data[k].Connect.Parameters = nil
	   			p.data2[k].Connect.Parameters = nil
	   		}

	   		 		// Connect
	   		   		// Delete
	   		   		// Get
	   		   		// Head
	   		   		// Options
	   		   		// Patch
	   		   		// Post
	   		   		// Put
	   		   		// Trace
	   	} 	*/

	// paths changelog
	changes, err := p.diff(p.data, p.data2)
	if err != nil {
		return err
	}

	// changelogs
	return p.handleChanges(changes)
}

func (p *pathsDiff) handleChanges(changes lib.Changelog) error {
	var err error

	for _, c := range changes {
		key := c.Path[0]
		lastPath := c.Path[len(c.Path)-1]
		penultPath := lastPath
		if len(c.Path) > 1 {
			penultPath = c.Path[len(c.Path)-2]
		}

		// this change is an array
		if penultPath != lastPath && model.IsArrayProperty(penultPath) {
			// lastPath -> array identifier
			// penultPath -> array property name

			// paths.servers
			if penultPath == "servers" {
				err = p.handleArrayChange(p.data[key].Servers, p.data2[key].Servers, c)
				if err != nil {
					return err
				}
				continue
			}

			// paths.parameters || paths.operation.parameters
			if penultPath == "parameters" {
				// paths.parameters
				if len(c.Path) == 3 {
					err = p.handleArrayChange(p.data[key].Parameters, p.data2[key].Parameters, c)
					if err != nil {
						return err
					}
					continue
				}
				// paths.operation.parameters
				if len(c.Path) > 3 {
					var data model.Array
					var data2 model.Array

					switch c.Path[1] {
					case "connect":
						data = p.data[key].Connect.Parameters
						data2 = p.data2[key].Connect.Parameters
					case "delete":
						data = p.data[key].Delete.Parameters
						data2 = p.data2[key].Delete.Parameters
					case "get":
						data = p.data[key].Get.Parameters
						data2 = p.data2[key].Get.Parameters
					case "head":
						data = p.data[key].Head.Parameters
						data2 = p.data2[key].Head.Parameters
					case "options":
						data = p.data[key].Options.Parameters
						data2 = p.data2[key].Options.Parameters
					case "patch":
						data = p.data[key].Patch.Parameters
						data2 = p.data2[key].Patch.Parameters
					case "post":
						data = p.data[key].Post.Parameters
						data2 = p.data2[key].Post.Parameters
					case "put":
						data = p.data[key].Put.Parameters
						data2 = p.data2[key].Put.Parameters
					case "trace":
						data = p.data[key].Trace.Parameters
						data2 = p.data2[key].Trace.Parameters

					}
					err = p.handleArrayChange(data, data2, c)
					if err != nil {
						return err
					}
					continue
				}
			}

		}

		path := fmt.Sprintf("%s.%s", p.key, lastPath)
		index := -1

		// data -> file1 -> base
		// data2 -> file2

		if c.Type == "create" || c.Type == "delete" {
			// create will display the path as the new element url value
			// url is the identifier for the servers array, let's get the index of the new element based on the last path
			// file1 is always the base file
			// creation is always from file2
			// deletion is always from file1

			var filePath string

			if c.Type == "create" {
				filePath = p.filePath2
				index, err = p.data2[key].Parameters.SearchByIdentifier(lastPath)
			} else {
				filePath = p.filePath
				index, err = p.data[key].Parameters.SearchByIdentifier(lastPath)
			}

			if err != nil {
				return err
			}
			if index != -1 {
				path = fmt.Sprintf("%s#%s.%d", filePath, p.key, index)

			}
		}

		// TODO: Find the source file/index of the updated element property
		// ISSUE: The identifier will always be present on both files, we need more info than just the identifier to find the source of the update
		if c.Type == "update" {
			// url is the identifier for the servers array, let's get the index of the new element based on the penult path
			// we have to figure out if it was updated from file1 or file2

			// file1
			//index, err := p.data.SearchByIdentifier(penultPath)
			if err != nil {
				return err
			}
			if index != -1 {
				path = fmt.Sprintf("%s#%s.%d.%s", p.filePath, p.key, index, lastPath)

			} else {
				// file2
				//index, err := p.data2.SearchByIdentifier(penultPath)
				if err != nil {
					return err
				}
				if index != -1 {
					path = fmt.Sprintf("%s#%s.%d.%s", p.filePath2, p.key, index, lastPath)
				}
			}
		}

		p.changelog = append(p.changelog,
			&changelog{
				Type: c.Type,
				Path: path,
				From: c.From,
				To:   c.To,
			},
		)
	}

	return nil
}

func (p *pathsDiff) handleOperation(ops, ops2 *model.Operation) error {
	opsChanges, err := p.diff(ops, ops2)
	if err != nil {
		return err
	}
	err = p.handleChanges(opsChanges)
	if err != nil {
		return err
	}

	return nil
}

func (p *pathsDiff) handleArray(params, params2 model.Array) error {
	paramsChanges, err := p.diff(params, params2)
	if err != nil {
		return err
	}
	err = p.handleArrayChanges(params, params2, paramsChanges)
	if err != nil {
		return err
	}

	return nil
}
