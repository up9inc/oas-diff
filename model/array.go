package model

type Array interface {
	SearchByIdentifier(identifier interface{}) (int, error)
}

func IsArrayProperty(p string) bool {
	for _, a := range [...]string{"servers", "parameters"} {
		if a == p {
			return true
		}
	}
	return false
}
