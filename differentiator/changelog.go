package differentiator

import (
	"strings"
	"time"
)

type ExecutionStatus struct {
	BaseFilePath   string                `json:"baseFile"`
	SecondFilePath string                `json:"secondFile"`
	StartTime      time.Time             `json:"startTime"`
	ExecutionTime  string                `json:"executionTime"`
	ExecutionFlags DifferentiatorOptions `json:"executionFlags"`
}
type ChangelogOutput struct {
	ExecutionStatus ExecutionStatus `json:"executionStatus"`
	Changelog       ChangeMap       `json:"changeLog"`
}

type ChangeMap map[string][]Changelog

type Changelog struct {
	Type       string      `json:"type"`
	Path       []string    `json:"path"`
	Identifier Identifier  `json:"identifier,omitempty"`
	From       interface{} `json:"from"`
	To         interface{} `json:"to"`
}

type ByType []Changelog

func (t ByType) Len() int           { return len(t) }
func (t ByType) Less(i, j int) bool { return t[i].Type < t[j].Type }
func (t ByType) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func NewChangelogOutput(startTime time.Time, baseFilePath, secondFilePath string, opts DifferentiatorOptions) *ChangelogOutput {
	return &ChangelogOutput{
		ExecutionStatus: ExecutionStatus{
			BaseFilePath:   baseFilePath,
			SecondFilePath: secondFilePath,
			StartTime:      startTime,
			ExecutionTime:  time.Since(startTime).String(),
			ExecutionFlags: opts,
		},
		Changelog: make(ChangeMap, 0),
	}
}

func (c ChangeMap) FilterByType(t string) ChangeMap {
	if len(t) == 0 {
		return c
	}

	filterType := strings.ToLower(t)
	filtered := make(ChangeMap, 0)

	for k, m := range c {
		for _, cc := range m {
			if cc.Type == filterType {
				filtered[k] = append(filtered[k], cc)
			}
		}
	}

	return filtered
}
