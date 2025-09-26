package alils

import (
	"fmt"
	"testing"
)

func TestGetLogStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogProject{
		Name:            "test",
		Endpoint:        "127.0.0.1:8888",
		AccessKeyID:     "access_key_id",
		AccessKeySecret: "access_key_secret",
	}
	s, err := p.GetLogStore("test")
	if err != nil {
		t.Errorf("failed to get logstore, err: %v", err)
		return
	}
	if s.Name != "test" {
		t.Errorf("failed to get logstore, name: %v", s.Name)
	}
}
