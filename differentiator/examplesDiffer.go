package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type examplesDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewExamplesDiffer(opts DifferentiatorOptions) *examplesDiffer {
	return &examplesDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (e *examplesDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Examples{}))
}

func (e *examplesDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if e.opts.Loose {
		aValue, aOk := a.Interface().(model.Examples)
		bValue, bOk := b.Interface().(model.Examples)

		if aOk && bOk {
			for ak, av := range aValue {
				for bk, bv := range bValue {
					// Ignore map key case sensitive
					if len(ak) > 0 && len(bk) > 0 && ak != bk && strings.EqualFold(ak, bk) {
						delete(aValue, ak)
						aValue[strings.ToLower(ak)] = av

						delete(bValue, bk)
						bValue[strings.ToLower(bk)] = bv
					}
				}
			}
		}
	}

	return e.differ.DiffMap(path, a, b)
}

func (e *examplesDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	e.DiffFunc = dfunc
}
