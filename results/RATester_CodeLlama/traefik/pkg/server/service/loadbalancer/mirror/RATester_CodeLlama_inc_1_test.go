package mirror

import (
	"fmt"
	"testing"
)

func TestInc_1(t *testing.T) {
	// setup types
	var m *Mirroring

	// setup tests
	tests := []struct {
		name string
		m    *Mirroring
	}{
		{
			"test inc",
			m,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.m.inc()
		})
	}
}
