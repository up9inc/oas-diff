package model

import file "github.com/up9inc/oas-diff/json"

type Model interface {
	// Each model struct must have its own parse logic
	Parse(file file.JsonFile) error
}