package alils

import (
	"fmt"
	"testing"
)

func TestCreateConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test_key_id",
		AccessKeySecret: "test_key_secret",
	}

	config := &LogConfig{
		Name:       "test_config",
		InputType:  "test_input_type",
		OutputType: "test_output_type",
		project:    project,
	}

	err := project.CreateConfig(config)
	if err != nil {
		t.Errorf("Failed to create config: %v", err)
	}
}
