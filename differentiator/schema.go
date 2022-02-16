package differentiator

import "github.com/up9inc/oas-diff/validator"

type schema struct {
	key                string
	properties         []string
	requiredProperties []string
}

func NewSchema(key string) *schema {
	return &schema{
		key: key,
	}
}

func (s *schema) Build(validator validator.Validator) error {
	var err error

	// properties
	s.properties, err = validator.GetSchemaPropertyFields(s.key)
	if err != nil {
		return err
	}

	// required properties
	s.requiredProperties, err = validator.GetSchemaPropertyRequiredFields(s.key)
	if err != nil {
		return err
	}

	return nil
}
