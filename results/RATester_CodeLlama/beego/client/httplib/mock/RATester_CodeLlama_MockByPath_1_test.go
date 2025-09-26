package mock

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestMockByPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MockResponseFilter{}
	path := "/path"
	resp := &http.Response{}
	err := errors.New("error")
	m.MockByPath(path, resp, err)
	if len(m.ms) != 1 {
		t.Fatalf("m.ms length is not 1")
	}
	if m.ms[0].cond.(*SimpleCondition).path != path {
		t.Fatalf("m.ms[0].cond.path is not %s", path)
	}
	if m.ms[0].resp != resp {
		t.Fatalf("m.ms[0].resp is not %v", resp)
	}
	if m.ms[0].err != err {
		t.Fatalf("m.ms[0].err is not %v", err)
	}
}
