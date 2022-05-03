package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type tagDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewTagDiffer(opts DifferentiatorOptions) *tagDiffer {
	return &tagDiffer{
		opts: opts,
	}
}

func (t *tagDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Tag{}))
}

func (t *tagDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if t.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.Tag)
		bValue, bOk := b.Interface().(model.Tag)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (t *tagDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	t.DiffFunc = dfunc
}
