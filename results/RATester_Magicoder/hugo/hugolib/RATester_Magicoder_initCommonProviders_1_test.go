package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestinitCommonProviders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ps := &pageState{
		pid: 1,
	}

	pp := pagePaths{
		outputFormats:        []page.OutputFormat{},
		firstOutputFormat:    page.OutputFormat{},
		targetPaths:          map[string]targetPathsHolder{},
		targetPathDescriptor: page.TargetPathDescriptor{},
	}

	err := ps.initCommonProviders(pp)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Add more test cases as needed
}
