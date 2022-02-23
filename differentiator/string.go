package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
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

func (differ *StringDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf((*string)(nil)).Elem())
}

func (differ *StringDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value) error {
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

	if differ.opts.Loose {
		if !strings.EqualFold(source, target) {
			cl.Add(lib.UPDATE, path, a.Interface(), b.Interface())
		}
	} else {
		if strings.Compare(source, target) != 0 {
			cl.Add(lib.UPDATE, path, a.Interface(), b.Interface())
		}
	}

	return nil
}

func (differ *StringDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	differ.DiffFunc = dfunc
}
