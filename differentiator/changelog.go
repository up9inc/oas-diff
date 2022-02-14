package differentiator

import (
	lib "github.com/r3labs/diff/v2"
)

type changelog struct {
	Key           string `json:"property"`
	lib.Changelog `json:"changelog"`
}

func NewChangelog(key string) *changelog {
	return &changelog{
		Key: key,
	}
}
