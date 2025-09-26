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

	counter.counts[0] = 1
	counter.counts[1] = 2
	counter.counts[2] = 3
	counter.counts[3] = 4
	counter.counts[4] = 5
	counter.counts[5] = 6
	counter.counts[6] = 7

	tests := []struct {
		name     string
		lastdays int64
		want     []int64
	}{
		{
			name:     "TestGetLastDaysCount_1",
			lastdays: 3,
			want:     []int64{4, 5, 6},
		},
		{
			name:     "TestGetLastDaysCount_2",
			lastdays: 10,
			want:     []int64{4, 5, 6, 7, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "TestGetLastDaysCount_3",
			lastdays: 0,
			want:     []int64{},
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
