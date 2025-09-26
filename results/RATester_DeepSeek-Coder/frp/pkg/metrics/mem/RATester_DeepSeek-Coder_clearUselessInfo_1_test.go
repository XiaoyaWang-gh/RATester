package mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestClearUselessInfo_1(t *testing.T) {
	type fields struct {
		info *ServerStatistics
		mu   sync.Mutex
	}
	type args struct {
		continuousOfflineDuration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  int
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
			got, got1 := m.clearUselessInfo(tt.args.continuousOfflineDuration)
			if got != tt.want {
				t.Errorf("serverMetrics.clearUselessInfo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("serverMetrics.clearUselessInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
