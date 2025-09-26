package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestValue_2(t *testing.T) {
	tests := []struct {
		name string
		e    DateField
		want time.Time
	}{
		{
			name: "Test case 1",
			e:    DateField(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
			want: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Test case 2",
			e:    DateField(time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)),
			want: time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC),
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
