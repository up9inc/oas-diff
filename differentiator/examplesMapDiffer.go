package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type examplesMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewExamplesMapDiffer(opts DifferentiatorOptions) *examplesMapDiffer {
	return &examplesMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (e *examplesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ExamplesMap{}))
}

func (e *examplesMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if e.opts.Loose {
		handleLooseMap[model.ExamplesMap](a, b)
	}

	return e.differ.DiffMap(path, a, b)
}

func (e *examplesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	e.DiffFunc = dfunc
}
