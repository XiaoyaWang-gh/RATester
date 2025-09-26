package mock

import (
	"fmt"
	"testing"
)

func TestMock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &OrmStub{}
	m := &Mock{}
	o.Mock(m)

	if len(o.ms) != 1 {
		t.Errorf("Expected 1 mock, got %d", len(o.ms))
	}

	if o.ms[0] != m {
		t.Errorf("Expected mock to be %v, got %v", m, o.ms[0])
	}
}
