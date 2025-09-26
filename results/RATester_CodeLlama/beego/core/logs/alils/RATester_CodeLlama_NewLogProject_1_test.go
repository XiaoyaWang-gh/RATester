package alils

import (
	"fmt"
	"testing"
)

func TestNewLogProject_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	endpoint := "127.0.0.1:8080"
	AccessKeyID := "AccessKeyID"
	accessKeySecret := "accessKeySecret"
	p, err := NewLogProject(name, endpoint, AccessKeyID, accessKeySecret)
	if err != nil {
		t.Errorf("NewLogProject error: %v", err)
	}
	if p.Name != name {
		t.Errorf("NewLogProject error: p.Name = %v, want %v", p.Name, name)
	}
	if p.Endpoint != endpoint {
		t.Errorf("NewLogProject error: p.Endpoint = %v, want %v", p.Endpoint, endpoint)
	}
	if p.AccessKeyID != AccessKeyID {
		t.Errorf("NewLogProject error: p.AccessKeyID = %v, want %v", p.AccessKeyID, AccessKeyID)
	}
	if p.AccessKeySecret != accessKeySecret {
		t.Errorf("NewLogProject error: p.AccessKeySecret = %v, want %v", p.AccessKeySecret, accessKeySecret)
	}
}
