package mock

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestNewMockByPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "path"
	resp := &http.Response{}
	err := errors.New("error")
	mock := NewMockByPath(path, resp, err)
	if mock.cond.(*SimpleCondition).path != path {
		t.Errorf("NewMockByPath() = %v, want %v", mock.cond.(*SimpleCondition).path, path)
	}
	if mock.resp != resp {
		t.Errorf("NewMockByPath() = %v, want %v", mock.resp, resp)
	}
	if mock.err != err {
		t.Errorf("NewMockByPath() = %v, want %v", mock.err, err)
	}
}
