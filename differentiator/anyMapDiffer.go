package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type anyMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewAnyMapDiffer(opts DifferentiatorOptions) *anyMapDiffer {
	return &anyMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (ad *anyMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.AnyMap{}))
}

func (ad *anyMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if ad.opts.Loose {
		aValue, aOk := a.Interface().(model.AnyMap)
		bValue, bOk := b.Interface().(model.AnyMap)

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

	return ad.differ.DiffMap(path, a, b)
}

func (ad *anyMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	ad.DiffFunc = dfunc
}
