package differentiator

import (
	"reflect"

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
	return p.differ.DiffSlice(path, a, b)
}

func (p *ParametersDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}
