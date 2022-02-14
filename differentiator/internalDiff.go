package differentiator

type internalDiff struct {
	key       string
	schema    *schema
	changelog *changelog
}

func NewInternalDiff(key string) *internalDiff {
	return &internalDiff{
		key:       key,
		schema:    NewSchema(key),
		changelog: NewChangelog(key),
	}
}
