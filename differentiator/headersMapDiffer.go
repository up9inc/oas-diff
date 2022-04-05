package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type headersMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewHeadersMapDiffer(opts DifferentiatorOptions) *headersMapDiffer {
	return &headersMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (h *headersMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.HeadersMap{}))
}

func (h *headersMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if h.opts.IgnoreExamples {
		ignoreExamplesFromMaps[model.HeadersMap](a, b)
	}

	if h.opts.Loose {
		handleLooseMap[model.HeadersMap](a, b)
	}

	return df(path, a, b, parent)
}

func (h *headersMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	h.DiffFunc = dfunc
}
