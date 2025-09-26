package page

import (
	"fmt"
	"testing"
)

func TestInit_45(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PageMatcherParamsConfig{}
	if err := p.init(); err != nil {
		t.Errorf("init() error = %v", err)
		return
	}
}
