package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type securitySchemeDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSecuritySchemeDiffer(opts DifferentiatorOptions) *securitySchemeDiffer {
	return &securitySchemeDiffer{
		opts: opts,
	}
}

func (s *securitySchemeDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.SecurityScheme{}))
}

func (s *securitySchemeDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.SecurityScheme)
		bValue, bOk := b.Interface().(model.SecurityScheme)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (s *securitySchemeDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
