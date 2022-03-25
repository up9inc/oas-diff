package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type requestBodiesMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewRequestBodiesMapDiffer(opts DifferentiatorOptions) *requestBodiesMapDiffer {
	return &requestBodiesMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (r *requestBodiesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.RequestBodiesMap{}))
}

func (r *requestBodiesMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if r.opts.Loose {
		aValue, aOk := a.Interface().(model.RequestBodiesMap)
		bValue, bOk := b.Interface().(model.RequestBodiesMap)

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

func (r *requestBodiesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	r.DiffFunc = dfunc
}
