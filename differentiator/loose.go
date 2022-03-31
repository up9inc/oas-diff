package differentiator

import (
	"reflect"
	"strings"

	"github.com/up9inc/oas-diff/model"
)

// TODO: HandleLooseArrays/Slices

func handleLooseMap[T model.MapsConstraint[V], V any](a, b reflect.Value) {
	aValue, aOk := a.Interface().(T)
	bValue, bOk := b.Interface().(T)

	if aOk && bOk {
		for ak, av := range aValue {
			for bk, bv := range bValue {
				// Ignore map key case sensitive
				if len(ak) > 0 && len(bk) > 0 && ak != bk && strings.EqualFold(ak, bk) {
					delete(aValue, ak)
					aValue[strings.ToLower(ak)] = av

					delete(bValue, bk)
					bValue[strings.ToLower(bk)] = bv
				}
			}
		}
	}
}

func ignoreExamplesFromMaps[T model.MapsConstraint[V], V model.ExamplesInterface](a, b reflect.Value) {
	aValue, aOk := a.Interface().(T)
	bValue, bOk := b.Interface().(T)

	if aOk {
		for _, av := range aValue {
			av.IgnoreExamples()
		}
	}

	if bOk {
		for _, bv := range bValue {
			bv.IgnoreExamples()
		}
	}
}

func ignoreExamplesFromSlices[T model.SlicesConstraint[V], V model.ExamplesInterface](a, b reflect.Value) {
	aValue, aOk := a.Interface().(T)
	bValue, bOk := b.Interface().(T)

	if aOk {
		for _, av := range aValue {
			av.IgnoreExamples()
		}
	}

	if bOk {
		for _, bv := range bValue {
			bv.IgnoreExamples()
		}
	}
}
