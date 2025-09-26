package alils

import (
	"fmt"
	"testing"
)

func TestGetAppliedConfigs_1(t *testing.T) {
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

	groupName := "test"

	confNames, err := p.GetAppliedConfigs(groupName)
	if err != nil {
		t.Errorf("failed to get applied configs, err: %v", err)
		return
	}

	if len(confNames) == 0 {
		t.Errorf("failed to get applied configs, confNames: %v", confNames)
		return
	}

	t.Logf("get applied configs, confNames: %v", confNames)
}
