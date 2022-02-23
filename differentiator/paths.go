package differentiator

import (
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

func (p *pathsDiff) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions) error {
	var err error

	// opts
	p.opts = opts

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

	// paths changelog
	changes, err := p.diff(p.data, p.data2)
	if err != nil {
		return err
	}

	// changelogs
	return p.handleChanges(changes)
}

func (p *pathsDiff) handleChanges(changes lib.Changelog) (err error) {
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

		// handle everything else
		p.internalDiff.handleChange(c)
	}

	return nil
}
