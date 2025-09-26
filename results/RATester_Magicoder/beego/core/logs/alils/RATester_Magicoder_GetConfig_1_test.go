package alils

import (
	"fmt"
	"testing"
)

func TestGetConfig_1(t *testing.T) {
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
	}

	err := project.CreateConfig(config)
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	got, err := project.GetConfig("test_config")
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}

	if got.Name != config.Name || got.InputType != config.InputType || got.OutputType != config.OutputType {
		t.Errorf("GetConfig() = %v, want %v", got, config)
	}

	err = project.DeleteConfig("test_config")
	if err != nil {
		t.Fatalf("Failed to delete config: %v", err)
	}
}
