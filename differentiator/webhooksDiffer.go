package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*webhooksDiffer)(nil)

type webhooksDiffer struct {
	*internalDiff
	data  model.Webhooks
	data2 model.Webhooks

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewWebhooksDiffer() *webhooksDiffer {
	return &webhooksDiffer{
		internalDiff: NewInternalDiff(model.OAS_WEBHOOKS_KEY),
		data:         model.Webhooks{},
		data2:        model.Webhooks{},
	}
}

func (w *webhooksDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
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

func (w *webhooksDiffer) handleChanges(changes lib.Changelog) (err error) {
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
			err = w.handleArrayChange(w.data[key].Servers, w.data2[key].Servers, c)
			if err != nil {
				return err
			}
			continue
		}

		// paths.parameters || paths.operation.parameters
		if isParametersArray {
			// paths.parameters
			if len(c.Path) == 3 {
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

func (w *webhooksDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Webhooks{}))
}

func (w *webhooksDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if w.opts.Loose {
		aValue, aOk := a.Interface().(model.Webhooks)
		bValue, bOk := b.Interface().(model.Webhooks)

		if aOk && bOk {
			for ak, av := range aValue {
				for bk, bv := range bValue {
					// Ignore map key case sensitive
					if len(ak) > 0 && len(bk) > 0 && ak != bk && strings.EqualFold(ak, bk) {
						delete(aValue, ak)
						aValue[strings.ToLower(ak)] = av

						delete(bValue, bk)
						bValue[strings.ToLower(bk)] = bv
					}
				}
			}
		}
	}

	return w.differ.DiffMap(path, a, b)
}

func (w *webhooksDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	w.DiffFunc = dfunc
}
