package util

func SliceElementExists[S ~[]V, V comparable](s S, v V) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}

	return false
}

func SliceFindElementIndex[S ~[]V, V comparable](s S, v V) int {
	if s == nil || len(s) == 0 {
		return -1
	}

	for si, sv := range s {
		if sv == v {
			return si
		}
	}

	return -1
}

func SliceElementAdd[S ~[]V, V comparable](s S, v V) S {
	return append(s, v)
}

func SliceElementAddUnique[S ~[]V, V comparable](s S, v V) S {
	if !SliceElementExists(s, v) {
		s = append(s, v)
	}

	return s
}

func SliceElementRemoveAtIndex[S ~[]V, V comparable](s S, i int) S {
	if len(s) == 0 {
		return s
	}

	// this will remove and change order, do not use it if the slice is ordered
	// Remove the element at index i from s
	s[i] = s[len(s)-1] // Copy last element to index i
	s = s[:len(s)-1]   // Truncate slice
	return s
}

func SliceElementRemoveOrderedAtIndex[S ~[]V, V comparable](s S, i int) S {
	if len(s) == 0 {
		return s
	}

	// Remove the element at index i from s
	copy(s[i:], s[i+1:]) // Shift a[i+1:] left one index
	s = s[:len(s)-1]     // Truncate slice
	return s
}
