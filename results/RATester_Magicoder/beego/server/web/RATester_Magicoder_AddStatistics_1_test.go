package web

import (
	"fmt"
	"testing"
	"time"
)

func TestAddStatistics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	urlMap := &URLMap{
		LengthLimit: 0,
		urlmap:      make(map[string]map[string]*Statistics),
	}

	urlMap.AddStatistics("GET", "/test", "TestController", time.Duration(1000000000))
	urlMap.AddStatistics("POST", "/test", "TestController", time.Duration(2000000000))
	urlMap.AddStatistics("GET", "/test", "TestController", time.Duration(3000000000))

	statistics := urlMap.urlmap["/test"]["GET"]
	if statistics.RequestNum != 2 {
		t.Errorf("Expected RequestNum to be 2, but got %d", statistics.RequestNum)
	}
	if statistics.MinTime != time.Duration(1000000000) {
		t.Errorf("Expected MinTime to be 1000000000, but got %d", statistics.MinTime)
	}
	if statistics.MaxTime != time.Duration(3000000000) {
		t.Errorf("Expected MaxTime to be 3000000000, but got %d", statistics.MaxTime)
	}
	if statistics.TotalTime != time.Duration(4000000000) {
		t.Errorf("Expected TotalTime to be 4000000000, but got %d", statistics.TotalTime)
	}
}
