package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestValue_10(t *testing.T) {
	tests := []struct {
		name string
		e    TimeField
		want time.Time
	}{
		{
			name: "Test case 1",
			e:    TimeField(time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC)),
			want: time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "Test case 2",
			e:    TimeField(time.Date(2022, 2, 1, 11, 0, 0, 0, time.UTC)),
			want: time.Date(2022, 2, 1, 11, 0, 0, 0, time.UTC),
		},
		{
			name: "Test case 3",
			e:    TimeField(time.Date(2022, 3, 1, 12, 0, 0, 0, time.UTC)),
			want: time.Date(2022, 3, 1, 12, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
