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
		t.Errorf("o.ms should have 1 element")
	}
}
