package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
)

func TestMatchCacheBuster_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var s BuildConfig
	var logger loggers.Logger
	var p string

	// When
	matcher, err := s.MatchCacheBuster(logger, p)

	// Then
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if matcher == nil {
		t.Errorf("expected matcher, got nil")
	}
}
