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

	p := &LogProject{
		Name:            "test",
		Endpoint:        "http://127.0.0.1:8080",
		AccessKeyID:     "accessKeyID",
		AccessKeySecret: "accessKeySecret",
	}

	c := &LogConfig{
		Name:         "test",
		InputType:    "test",
		InputDetail:  InputDetail{},
		OutputType:   "test",
		OutputDetail: OutputDetail{},
		project:      p,
	}

	err := p.CreateConfig(c)
	if err != nil {
		t.Errorf("failed to create config: %v", err)
	}
}
