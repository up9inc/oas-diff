package model

type Array interface {
	GetName() string
	GetIdentifierName() string
	SearchByIdentifier(identifier interface{}) (int, error)
	FilterIdentifiers() []*ArrayIdentifierFilter
}

type ArrayIdentifierFilter struct {
	Name  string
	Index int
}
