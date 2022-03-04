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

// TODO: Header identification for the loose flag logic
// TODO: Test if response Header is catched here
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
			aIds := aValue.FilterIdentifiers()
			bIds := bValue.FilterIdentifiers()

			for _, a := range aIds {
				for _, b := range bIds {
					if a.Name != b.Name && strings.EqualFold(a.Name, b.Name) {
						// Only remove the data for headers when we want to completly ignore it and all its sub data
						//aValue[a.index] = nil
						//bValue[b.index] = nil

						// we don't want this case sensitive identifier comparison
						// set lower case for both identifiers and keep comparing
						aValue[a.Index].Name = strings.ToLower(aValue[a.Index].Name)
						bValue[b.Index].Name = strings.ToLower(bValue[b.Index].Name)
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
