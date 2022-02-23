package model

type Array interface {
	SearchByIdentifier(identifier interface{}) (int, error)
}

func IsArrayProperty(property string) bool {
	for _, a := range [...]string{"servers", "parameters"} {
		if a == property {
			return true
		}
	}
	return false
}
