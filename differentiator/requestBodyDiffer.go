package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type requestBodyDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewRequestBodyDiffer(opts DifferentiatorOptions) *requestBodyDiffer {
	return &requestBodyDiffer{
		opts: opts,
	}
}

func (r *requestBodyDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.RequestBody{}))
}

func (r *requestBodyDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if r.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.RequestBody)
		bValue, bOk := b.Interface().(model.RequestBody)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (r *requestBodyDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	r.DiffFunc = dfunc
}
