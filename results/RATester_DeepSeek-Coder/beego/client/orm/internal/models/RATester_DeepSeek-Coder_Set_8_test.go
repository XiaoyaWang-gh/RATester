package models

import (
	"fmt"
	"testing"
)

func TestSet_8(t *testing.T) {
	tests := []struct {
		name string
		e    FloatField
		d    float64
		want FloatField
	}{
		{
			name: "Test case 1",
			e:    FloatField(1.0),
			d:    2.0,
			want: FloatField(2.0),
		},
		{
			name: "Test case 2",
			e:    FloatField(3.0),
			d:    4.0,
			want: FloatField(4.0),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := &tt.e
			e.Set(tt.d)
			if *e != tt.want {
				t.Errorf("FloatField.Set() = %v, want %v", *e, tt.want)
			}
		})
	}
}
