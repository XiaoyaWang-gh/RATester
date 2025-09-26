package envvar

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNewEnvVar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	cfg := Config{
		ExportVars: map[string]string{
			"key1": "val1",
			"key2": "val2",
		},
		ExcludeVars: map[string]string{
			"key3": "val3",
			"key4": "val4",
		},
	}

	// when
	envVar := newEnvVar(cfg)

	// then
	assert.Equal(t, map[string]string{
		"key1": "val1",
		"key2": "val2",
	}, envVar.Vars)
}
