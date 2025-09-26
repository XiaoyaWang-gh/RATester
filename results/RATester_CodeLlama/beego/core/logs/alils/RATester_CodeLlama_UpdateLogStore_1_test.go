package alils

import (
	"fmt"
	"testing"
)

func TestUpdateLogStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogProject{
		Name:            "test",
		Endpoint:        "http://127.0.0.1:8888",
		AccessKeyID:     "accessKeyID",
		AccessKeySecret: "accessKeySecret",
	}

	err := p.UpdateLogStore("test", 1, 1)
	if err != nil {
		t.Errorf("UpdateLogStore failed, err: %v", err)
	}
}
