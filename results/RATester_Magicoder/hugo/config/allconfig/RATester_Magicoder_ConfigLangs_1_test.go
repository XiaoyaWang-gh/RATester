package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestConfigLangs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configs := Configs{
		configLangs: []config.AllProvider{
			// Initialize with some mock data
		},
	}

	result := configs.ConfigLangs()

	if len(result) != len(configs.configLangs) {
		t.Errorf("Expected length of result to be %d, but got %d", len(configs.configLangs), len(result))
	}

	for i, lang := range result {
		if lang != configs.configLangs[i] {
			t.Errorf("Expected lang at index %d to be %v, but got %v", i, configs.configLangs[i], lang)
		}
	}
}
