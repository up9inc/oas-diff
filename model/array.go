package model

type Array interface {
	SearchByIdentifier(identifier interface{}) (int, interface{}, error)
}
