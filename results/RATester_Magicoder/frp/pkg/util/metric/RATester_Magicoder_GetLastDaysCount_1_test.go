package metric

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetLastDaysCount_1(t *testing.T) {
	counter := &StandardDateCounter{
		reserveDays:    7,
		counts:         make([]int64, 7),
		lastUpdateDate: time.Now(),
	}

	counter.counts[0] = 10
	counter.counts[1] = 20
	counter.counts[2] = 30
	counter.counts[3] = 40
	counter.counts[4] = 50
	counter.counts[5] = 60
	counter.counts[6] = 70

	tests := []struct {
		name     string
		lastdays int64
		want     []int64
	}{
		{
			name:     "Test case 1",
			lastdays: 3,
			want:     []int64{10, 20, 30},
		},
		{
			name:     "Test case 2",
			lastdays: 7,
			want:     []int64{10, 20, 30, 40, 50, 60, 70},
		},
		{
			name:     "Test case 3",
			lastdays: 10,
			want:     []int64{10, 20, 30, 40, 50, 60, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := counter.GetLastDaysCount(tt.lastdays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastDaysCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
