package mock

import (
	"fmt"
	"testing"
)

func TestWithPathReg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pathReg := "pathReg"
	sc := &SimpleCondition{}
	WithPathReg(pathReg)(sc)
	if sc.pathReg != pathReg {
		t.Errorf("WithPathReg() failed")
	}
}
