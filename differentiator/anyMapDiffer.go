package differentiator

import (
	"reflect"

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

func (ad *anyMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if ad.opts.Loose {
		handleLooseMap[model.AnyMap](a, b)
	}

	return df(path, a, b, parent)
}

func (ad *anyMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	ad.DiffFunc = dfunc
}
