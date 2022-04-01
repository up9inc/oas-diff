package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type parameterDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewParameterDiffer(opts DifferentiatorOptions) *parameterDiffer {
	return &parameterDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (p *parameterDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Parameter{}))
}

func (p *parameterDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if p.opts.Loose {
		aValue, aOk := a.Interface().(model.Parameter)
		bValue, bOk := b.Interface().(model.Parameter)

		// ignore parameters header taht starts with x- or is an user-agent
		if aOk {
			if aValue.IsHeader() && aValue.IsIgnoredWhenLoose() {
				return nil
			}
		}

		if bOk {
			if bValue.IsHeader() && bValue.IsIgnoredWhenLoose() {
				return nil
			}
		}
	}

	return p.differ.DiffStruct(path, a, b)
}

func (p *parameterDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}
