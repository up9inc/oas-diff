package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type requestBodiesMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewRequestBodiesMapDiffer(opts DifferentiatorOptions) *requestBodiesMapDiffer {
	return &requestBodiesMapDiffer{
		opts: opts,
	}
}

func (r *requestBodiesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.RequestBodiesMap{}))
}

func (r *requestBodiesMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if r.opts.IgnoreDescriptions {
		ignoreDescriptionsFromMaps[model.RequestBodiesMap](a, b)
	}

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

	return df(path, a, b, parent)
}

func (r *requestBodiesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	r.DiffFunc = dfunc
}
