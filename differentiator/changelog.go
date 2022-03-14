package differentiator

import "time"

type changeMap map[string][]*changelog

type executionStatus struct {
	BaseFilePath   string                `json:"base-file"`
	SecondFilePath string                `json:"second-file"`
	StartTime      string                `json:"start-time"`
	ExecutionTime  string                `json:"execution-time"`
	ExecutionFlags DifferentiatorOptions `json:"execution-flags"`
}

type changelogOutput struct {
	ExecutionStatus executionStatus `json:"execution-status"`
	Changelog       changeMap       `json:"changelog"`
}

type changelog struct {
	Type       string      `json:"type"`
	Path       []string    `json:"path"`
	Identifier Identifier  `json:"identifier,omitempty"`
	From       interface{} `json:"from"`
	To         interface{} `json:"to"`
}

func NewChangelogOutput(startTime time.Time, baseFilePath, secondFilePath string, opts DifferentiatorOptions) *changelogOutput {
	return &changelogOutput{
		ExecutionStatus: executionStatus{
			BaseFilePath:   baseFilePath,
			SecondFilePath: secondFilePath,
			StartTime:      startTime.Format(time.StampMilli),
			ExecutionTime:  time.Since(startTime).String(),
			ExecutionFlags: opts,
		},
		Changelog: make(changeMap, 0),
	}
}
