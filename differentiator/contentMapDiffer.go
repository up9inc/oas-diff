package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type contentMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewContentMapDiffer(opts DifferentiatorOptions) *contentMapDiffer {
	return &contentMapDiffer{
		opts: opts,
	}
}

func (c *contentMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ContentMap{}))
}

func (c *contentMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if c.opts.IgnoreExamples {
		ignoreExamplesFromMaps[model.ContentMap](a, b)
	}

	if c.opts.Loose {
		handleLooseMap[model.ContentMap](a, b)
	}

	return df(path, a, b, parent)
}

func (c *contentMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	c.DiffFunc = dfunc
}
