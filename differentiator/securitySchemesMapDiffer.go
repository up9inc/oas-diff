package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type securitySchemesMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSecuritySchemesMapDiffer(opts DifferentiatorOptions) *securitySchemesMapDiffer {
	return &securitySchemesMapDiffer{
		opts: opts,
	}
}

func (s *securitySchemesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.SecuritySchemesMap{}))
}

func (s *securitySchemesMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.Loose {
		handleLooseMap[model.SecuritySchemesMap](a, b)
	}

	return df(path, a, b, parent)
}

func (s *securitySchemesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
