package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*pathsMapDiffer)(nil)

type pathsMapDiffer struct {
	*internalDiff
	data  model.PathsMap
	data2 model.PathsMap

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewPathsMapDiffer() *pathsMapDiffer {
	return &pathsMapDiffer{
		internalDiff: NewInternalDiff(model.OAS_PATHS_KEY),
		data:         model.PathsMap{},
		data2:        model.PathsMap{},
	}
}

func (p *pathsMapDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	p.opts = opts

	// differ
	p.differ = differ

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
	changes, err := p.differ.Diff(p.data, p.data2)
	if err != nil {
		return err
	}

	// changelogs
	return p.handleChanges(changes)
}

func (p *pathsMapDiffer) handleChanges(changes lib.Changelog) (err error) {
	for _, c := range changes {
		key := c.Path[0]

		var isServersArray bool
		var isParametersArray bool

		// Find array properties related to paths model
		serversName := model.Servers{}.GetName()
		parametersName := model.Parameters{}.GetName()

		for _, path := range c.Path {
			switch path {
			case serversName:
				isServersArray = true
			case parametersName:
				isParametersArray = true
			}
		}

		// paths.servers
		if isServersArray {
			err = p.handleArrayChange(p.data[key].Servers, p.data2[key].Servers, c)
			if err != nil {
				return err
			}
			continue
		}

		// paths.parameters || paths.operation.parameters
		if isParametersArray {
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

		// handle everything else
		p.internalDiff.handleChange(c)
	}

	return nil
}

func (p *pathsMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.PathsMap{}))
}

func (p *pathsMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if p.opts.Loose {
		handleLooseMap[model.PathsMap](a, b)
	}

	return p.differ.DiffMap(path, a, b)
}

func (p *pathsMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}
