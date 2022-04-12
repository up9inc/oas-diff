package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type encodingMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewEncodingMapDiffer(opts DifferentiatorOptions) *encodingMapDiffer {
	return &encodingMapDiffer{
		opts: opts,
	}
}

func (e *encodingMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.EncodingMap{}))
}

func (e *encodingMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if e.opts.Loose {
		handleLooseMap[model.EncodingMap](a, b)
	}

	return df(path, a, b, parent)
}

func (e *encodingMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	e.DiffFunc = dfunc
}
