package differentiator

type changeMap map[string][]*changelog

type changelog struct {
	Type string      `json:"type"`
	Path []string    `json:"path"`
	From interface{} `json:"from"`
	To   interface{} `json:"to"`
}

func NewChangeMap() changeMap {
	return make(changeMap, 0)
}
