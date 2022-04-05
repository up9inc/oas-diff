package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type linksMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewLinksMapDiffer(opts DifferentiatorOptions) *linksMapDiffer {
	return &linksMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (l *linksMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.LinksMap{}))
}

func (l *linksMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if l.opts.Loose {
		handleLooseMap[model.LinksMap](a, b)
	}

	return df(path, a, b, parent)
}

func (l *linksMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	l.DiffFunc = dfunc
}
