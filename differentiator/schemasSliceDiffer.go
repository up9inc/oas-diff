package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type schemasSliceDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSchemasSliceDiffer(opts DifferentiatorOptions) *schemasSliceDiffer {
	return &schemasSliceDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (s *schemasSliceDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.SchemasSlice{}))
}

func (s *schemasSliceDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	/* 	aValue, aOk := a.Interface().(model.SchemasSlice)
	   	bValue, bOk := b.Interface().(model.SchemasSlice)

	   	if aOk {
	   		fmt.Println(aValue)
	   	}

	   	if bOk {
	   		fmt.Println(bValue)
	   	} */

	return df(path, a, b, parent)
}

func (s *schemasSliceDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
