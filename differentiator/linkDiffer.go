package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type linkDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewLinkDiffer(opts DifferentiatorOptions) *linkDiffer {
	return &linkDiffer{
		opts: opts,
	}
}

func (l *linkDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Link{}))
}

func (l *linkDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if l.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.Link)
		bValue, bOk := b.Interface().(model.Link)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (l *linkDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	l.DiffFunc = dfunc
}
