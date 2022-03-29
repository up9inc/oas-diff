package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type parametersMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewParametersMapDiffer(opts DifferentiatorOptions) *parametersMapDiffer {
	return &parametersMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (p *parametersMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ParametersMap{}))
}

func (p *parametersMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if p.opts.Loose {
		handleLooseMap[model.ParametersMap](a, b)
	}

	return p.differ.DiffMap(path, a, b)
}

func (p *parametersMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}
