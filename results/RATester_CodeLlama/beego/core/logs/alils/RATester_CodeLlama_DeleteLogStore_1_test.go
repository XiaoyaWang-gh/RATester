package alils

import (
	"fmt"
	"testing"
)

func TestDeleteLogStore_1(t *testing.T) {
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
	err := p.DeleteLogStore("test")
	if err != nil {
		t.Errorf("failed to delete logstore")
	}
}
