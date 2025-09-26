package alils

import (
	"fmt"
	"testing"
)

func TestDeleteConfig_1(t *testing.T) {
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
	err := p.DeleteConfig("test")
	if err != nil {
		t.Errorf("failed to delete config")
	}
}
