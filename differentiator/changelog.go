package differentiator

import (
	lib "github.com/r3labs/diff/v2"
)

type changelog struct {
	key string
	lib.Changelog
}

func NewChangelog(key string) *changelog {
	return &changelog{
		key: key,
	}
}
