package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type encodingMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewEncodingMapDiffer(opts DifferentiatorOptions) *encodingMapDiffer {
	return &encodingMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (e *encodingMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.EncodingMap{}))
}

func (e *encodingMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if e.opts.Loose {
		aValue, aOk := a.Interface().(model.EncodingMap)
		bValue, bOk := b.Interface().(model.EncodingMap)

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

func (e *encodingMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	e.DiffFunc = dfunc
}
