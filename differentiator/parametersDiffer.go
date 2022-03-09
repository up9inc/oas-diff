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
			for ai, a := range aValue {
				for bi, b := range bValue {
					if a != nil && b != nil {
						// Ignore array Identifier case sensitive
						if len(a.Name) > 0 && len(b.Name) > 0 && a.Name != b.Name && strings.EqualFold(a.Name, b.Name) {
							// we don't want this case sensitive identifier comparison
							// set lower case for both identifiers and keep comparing
							aValue[ai].Name = strings.ToLower(aValue[ai].Name)
							bValue[bi].Name = strings.ToLower(bValue[bi].Name)
						}

						// Headers filter
						// TODO: Support reggex value
						if a.IsHeader() {
							// starts with x-
							if a.IsIgnoredWhenLoose() {
								aValue[ai] = nil
							}
						}
						if b.IsHeader() {
							if b.IsIgnoredWhenLoose() {
								bValue[bi] = nil
							}
						}
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
