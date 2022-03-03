package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type ParametersDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewParameterDiffer(opts DifferentiatorOptions) *ParametersDiffer {
	return &ParametersDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (p *ParametersDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Parameters{}))
}

func (p *ParametersDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if p.opts.Loose {
		if a.Kind() == reflect.Invalid {
			cl.Add(lib.CREATE, path, nil, lib.ExportInterface(b))
			return nil
		}

		if b.Kind() == reflect.Invalid {
			cl.Add(lib.DELETE, path, lib.ExportInterface(a), nil)
			return nil
		}

		if a.Kind() != b.Kind() {
			return lib.ErrTypeMismatch
		}

		aValue, aOk := a.Interface().(model.Parameters)
		bValue, bOk := b.Interface().(model.Parameters)

		if aOk && bOk {
			var aIds, bIds []struct {
				name  string
				index int
			}

			aIds = p.getParametersIdentifiers(aValue)
			bIds = p.getParametersIdentifiers(bValue)

			for _, a := range aIds {
				for _, b := range bIds {
					if a.name != b.name && strings.EqualFold(a.name, b.name) {
						// Only remove the data for headers when we want to completly ignore it and all its sub data
						//aValue[a.index] = nil
						//bValue[b.index] = nil

						// we don't want this case sensitive identifier comparison
						// set lower case for both identifiers and keep comparing
						aValue[a.index].Name = strings.ToLower(aValue[a.index].Name)
						bValue[b.index].Name = strings.ToLower(bValue[b.index].Name)
					}
				}
			}
		}
	}

	return p.differ.DiffSlice(path, a, b)
}

func (p *ParametersDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}

func (p *ParametersDiffer) getParametersIdentifiers(params model.Parameters) []struct {
	name  string
	index int
} {
	var result []struct {
		name  string
		index int
	}
	for aI, aP := range params {
		// name is the identifier
		if len(aP.Name) > 0 {
			result = append(result, struct {
				name  string
				index int
			}{
				name:  aP.Name,
				index: aI,
			})
		}
	}

	return result
}
