package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type operationDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewOperationDiffer(opts DifferentiatorOptions) *operationDiffer {
	return &operationDiffer{
		opts: opts,
	}
}

func (o *operationDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Operation{}))
}

func (o *operationDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if o.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.Operation)
		bValue, bOk := b.Interface().(model.Operation)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (o *operationDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	o.DiffFunc = dfunc
}
