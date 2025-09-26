package alils

import (
	"fmt"
	"testing"
)

func TestUpdateConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogProject{
		Name:            "test",
		Endpoint:        "127.0.0.1:8080",
		AccessKeyID:     "access_key_id",
		AccessKeySecret: "access_key_secret",
	}

	c := &LogConfig{
		Name:         "test",
		InputType:    "test",
		InputDetail:  InputDetail{},
		OutputType:   "test",
		OutputDetail: OutputDetail{},
		project:      p,
	}

	err := p.UpdateConfig(c)
	if err != nil {
		t.Errorf("failed to update config")
	}
}
