package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type callbacksMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewCallbacksMapDiffer(opts DifferentiatorOptions) *callbacksMapDiffer {
	return &callbacksMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (c *callbacksMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.CallbacksMap{}))
}

func (c *callbacksMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if c.opts.Loose {
		handleLooseMap[model.CallbacksMap](a, b)
	}

	return c.differ.DiffMap(path, a, b)
}

func (c *callbacksMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	c.DiffFunc = dfunc
}
