package debug

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestStop_2(t *testing.T) {
	type fields struct {
		start    time.Time
		elapsed  time.Duration
		stopOnce sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestStop",
			fields: fields{
				start:    time.Now(),
				stopOnce: sync.Once{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			timer := &timer{
				start:    tt.fields.start,
				elapsed:  tt.fields.elapsed,
				stopOnce: tt.fields.stopOnce,
			}
			if got := timer.Stop(); got != tt.want {
				t.Errorf("timer.Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}
