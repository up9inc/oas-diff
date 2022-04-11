package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v3"
)

type StringDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewStringDiffer(opts DifferentiatorOptions) *StringDiffer {
	return &StringDiffer{
		opts: opts,
	}
}

func (s *StringDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf((*string)(nil)).Elem())
}

func (s *StringDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreDescriptions && s.IsDescription(path) {
		return nil
	}

	// loose flag logic
	if s.opts.Loose {
		if a.Kind() == reflect.Invalid {
			cl.Add(lib.CREATE, path, nil, b.Interface())
			return nil
		}

		if b.Kind() == reflect.Invalid {
			cl.Add(lib.DELETE, path, a.Interface(), nil)
			return nil
		}

		var source, target string
		source, _ = a.Interface().(string)
		target, _ = b.Interface().(string)

		// TODO: Ignore EqualFold when the property is regex?
		if !strings.EqualFold(source, target) {
			cl.Add(lib.UPDATE, path, a.Interface(), b.Interface())
		}
		return nil
	}

	return df(path, a, b, parent)
}

func (s *StringDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}

func (s *StringDiffer) IsDescription(path []string) bool {
	if len(path) == 0 {
		return false
	}
	return path[len(path)-1] == "description"
}
