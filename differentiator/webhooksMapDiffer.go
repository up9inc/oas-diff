package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*webhooksMapDiffer)(nil)

type webhooksMapDiffer struct {
	*internalDiff
	data  model.WebhooksMap
	data2 model.WebhooksMap

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewWebhooksMapDiffer() *webhooksMapDiffer {
	return &webhooksMapDiffer{
		internalDiff: NewInternalDiff(model.OAS_WEBHOOKS_KEY),
		data:         model.WebhooksMap{},
		data2:        model.WebhooksMap{},
	}
}

func (w *webhooksMapDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	w.opts = opts

	// differ
	w.differ = differ

	// schema
	err = w.schema.Build(validator)
	if err != nil {
		return err
	}

	// webhooks1
	w.filePath = jsonFile.GetPath()
	err = w.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// webhooks2
	w.filePath2 = jsonFile2.GetPath()
	err = w.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// webhooks changelog
	changes, err := w.differ.Diff(w.data, w.data2)
	if err != nil {
		return err
	}

	// changelogs
	return w.handleChanges(changes)
}

func (w *webhooksMapDiffer) handleChanges(changes lib.Changelog) (err error) {
	for _, c := range changes {
		if len(c.Path) == 0 {
			w.internalDiff.handleChange(c)
			continue
		}

		key := c.Path[0]

		var isServersArray bool
		var isParametersArray bool

		// Find array properties related to paths model
		serversName := model.Servers{}.GetName()
		parametersName := model.Parameters{}.GetName()

		// get the latest identifier only, if multiple provided
		for i := len(c.Path) - 1; i >= 0; i-- {
			switch c.Path[i] {
			case serversName:
				isServersArray = true
			case parametersName:
				isParametersArray = true
			}
			if isServersArray || isParametersArray {
				break
			}
		}

		// paths.servers
		if isServersArray {
			err = w.handleArrayChange(w.data[key].Servers, w.data2[key].Servers, c)
			if err != nil {
				return err
			}
			continue
		}

		// paths.parameters || paths.operation.parameters
		if isParametersArray {
			// paths.parameters
			if len(c.Path) > 1 && c.Path[1] == parametersName {
				err = w.handleArrayChange(w.data[key].Parameters, w.data2[key].Parameters, c)
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
					data = w.data[key].Connect.Parameters
					data2 = w.data2[key].Connect.Parameters
				case "delete":
					data = w.data[key].Delete.Parameters
					data2 = w.data2[key].Delete.Parameters
				case "get":
					data = w.data[key].Get.Parameters
					data2 = w.data2[key].Get.Parameters
				case "head":
					data = w.data[key].Head.Parameters
					data2 = w.data2[key].Head.Parameters
				case "options":
					data = w.data[key].Options.Parameters
					data2 = w.data2[key].Options.Parameters
				case "patch":
					data = w.data[key].Patch.Parameters
					data2 = w.data2[key].Patch.Parameters
				case "post":
					data = w.data[key].Post.Parameters
					data2 = w.data2[key].Post.Parameters
				case "put":
					data = w.data[key].Put.Parameters
					data2 = w.data2[key].Put.Parameters
				case "trace":
					data = w.data[key].Trace.Parameters
					data2 = w.data2[key].Trace.Parameters

				}
				err = w.handleArrayChange(data, data2, c)
				if err != nil {
					return err
				}
				continue
			}
		}

		// handle everything else
		w.internalDiff.handleChange(c)
	}

	return nil
}

func (w *webhooksMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.WebhooksMap{}))
}

func (w *webhooksMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if w.opts.Loose {
		handleLooseMap[model.WebhooksMap](a, b)
	}

	return df(path, a, b, parent)
}

func (w *webhooksMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	w.DiffFunc = dfunc
}
