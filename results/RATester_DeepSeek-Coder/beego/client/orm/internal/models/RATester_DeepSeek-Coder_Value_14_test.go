package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestValue_14(t *testing.T) {
	tests := []struct {
		name string
		e    DateTimeField
		want time.Time
	}{
		{
			name: "Test 1",
			e:    DateTimeField(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)),
			want: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Test 2",
			e:    DateTimeField(time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)),
			want: time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			name: "Test 3",
			e:    DateTimeField(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)),
			want: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
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
				t.Errorf("DateTimeField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
