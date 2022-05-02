package differentiator

import (
	"reflect"

	"github.com/up9inc/oas-diff/model"
)

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

func ignoreDescriptionsFromMaps[T model.MapsConstraint[V], V model.DescriptionsInterface](a, b reflect.Value) {
	aValue, aOk := a.Interface().(T)
	bValue, bOk := b.Interface().(T)

	if aOk {
		for _, av := range aValue {
			av.IgnoreDescriptions()
		}
	}

	if bOk {
		for _, bv := range bValue {
			bv.IgnoreDescriptions()
		}
	}
}

func ignoreDescriptionsFromSlices[T model.SlicesConstraint[V], V model.DescriptionsInterface](a, b reflect.Value) {
	aValue, aOk := a.Interface().(T)
	bValue, bOk := b.Interface().(T)

	if aOk {
		for _, av := range aValue {
			av.IgnoreDescriptions()
		}
	}

	if bOk {
		for _, bv := range bValue {
			bv.IgnoreDescriptions()
		}
	}
}
