package mem

import (
	"fmt"
	"sync"
	"testing"
)

func TestRun_12(t *testing.T) {
	type fields struct {
		info *ServerStatistics
		mu   sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &serverMetrics{
				info: tt.fields.info,
				mu:   tt.fields.mu,
			}
			m.run()
		})
	}
}
