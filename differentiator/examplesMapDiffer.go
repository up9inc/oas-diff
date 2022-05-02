package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type examplesMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewExamplesMapDiffer(opts DifferentiatorOptions) *examplesMapDiffer {
	return &examplesMapDiffer{
		opts: opts,
	}
}

func (e *examplesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ExamplesMap{}))
}

func (e *examplesMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if e.opts.IgnoreDescriptions {
		ignoreDescriptionsFromMaps[model.ExamplesMap](a, b)
	}

	if e.opts.IgnoreExamples {
		return nil
	}

	if e.opts.Loose {
		handleLooseMap[model.ExamplesMap](a, b)
	}

	return df(path, a, b, parent)
}

func (e *examplesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	e.DiffFunc = dfunc
}
