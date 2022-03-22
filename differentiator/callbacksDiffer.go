package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type callbacksDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewCallbacksDiffer(opts DifferentiatorOptions) *callbacksDiffer {
	return &callbacksDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (c *callbacksDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Callbacks{}))
}

func (c *callbacksDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if c.opts.Loose {
		aValue, aOk := a.Interface().(model.Callbacks)
		bValue, bOk := b.Interface().(model.Callbacks)

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

	return c.differ.DiffMap(path, a, b)
}

func (c *callbacksDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	c.DiffFunc = dfunc
}
