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

	name := "test_project"
	endpoint := "http://test.com"
	AccessKeyID := "test_id"
	accessKeySecret := "test_secret"

	project, err := NewLogProject(name, endpoint, AccessKeyID, accessKeySecret)

	if err != nil {
		t.Errorf("NewLogProject() error = %v", err)
		return
	}

	if project.Name != name {
		t.Errorf("NewLogProject() name = %v, want %v", project.Name, name)
	}

	if project.Endpoint != endpoint {
		t.Errorf("NewLogProject() endpoint = %v, want %v", project.Endpoint, endpoint)
	}

	if project.AccessKeyID != AccessKeyID {
		t.Errorf("NewLogProject() AccessKeyID = %v, want %v", project.AccessKeyID, AccessKeyID)
	}

	if project.AccessKeySecret != accessKeySecret {
		t.Errorf("NewLogProject() accessKeySecret = %v, want %v", project.AccessKeySecret, accessKeySecret)
	}
}
