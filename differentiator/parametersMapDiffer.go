package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type parametersMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewParametersMapDiffer(opts DifferentiatorOptions) *parametersMapDiffer {
	return &parametersMapDiffer{
		opts: opts,
	}
}

func (p *parametersMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ParametersMap{}))
}

func (p *parametersMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if p.opts.IgnoreDescriptions {
		ignoreDescriptionsFromMaps[model.ParametersMap](a, b)
	}

	if p.opts.IgnoreExamples {
		ignoreExamplesFromMaps[model.ParametersMap](a, b)
	}

	if p.opts.Loose {
		handleLooseMap[model.ParametersMap](a, b)
	}

	return df(path, a, b, parent)
}

func (p *parametersMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}
