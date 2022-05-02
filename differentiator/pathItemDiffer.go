package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type pathItemDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewPathItemDiffer(opts DifferentiatorOptions) *pathItemDiffer {
	return &pathItemDiffer{
		opts: opts,
	}
}

func (p *pathItemDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.PathItem{}))
}

func (p *pathItemDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if p.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.PathItem)
		bValue, bOk := b.Interface().(model.PathItem)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (p *pathItemDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	p.DiffFunc = dfunc
}
