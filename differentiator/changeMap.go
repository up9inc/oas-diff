package differentiator

type changeMap map[string][]*changelog

type changelog struct {
	Type       string      `json:"type"`
	Path       []string    `json:"path"`
	Identifier Identifier  `json:"identifier,omitempty"`
	From       interface{} `json:"from"`
	To         interface{} `json:"to"`
}

func NewChangeMap() changeMap {
	return make(changeMap, 0)
}
