package prometheus

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestregisterBuildInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registerBuildInfo()

	// Check if the buildInfo metric is registered
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Errorf("Failed to gather metrics: %v", err)
	}

	found := false
	for _, mf := range metricFamilies {
		if *mf.Name == "beego" && *mf.Help == "The building information" {
			found = true
			break
		}
	}

	if !found {
		t.Error("buildInfo metric not found")
	}
}
