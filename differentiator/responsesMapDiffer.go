package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type responsesMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewResponsesMapDiffer(opts DifferentiatorOptions) *responsesMapDiffer {
	return &responsesMapDiffer{
		opts: opts,
	}
}

func (r *responsesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ResponsesMap{}))
}

func (r *responsesMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if r.opts.Loose {
		handleLooseMap[model.ResponsesMap](a, b)
	}

	return df(path, a, b, parent)
}

func (r *responsesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	r.DiffFunc = dfunc
}
