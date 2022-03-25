package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type responsesMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewResponsesMapDiffer(opts DifferentiatorOptions) *responsesMapDiffer {
	return &responsesMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (r *responsesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ResponsesMap{}))
}

func (r *responsesMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if r.opts.Loose {
		aValue, aOk := a.Interface().(model.ResponsesMap)
		bValue, bOk := b.Interface().(model.ResponsesMap)

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

	return r.differ.DiffMap(path, a, b)
}

func (r *responsesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	r.DiffFunc = dfunc
}
