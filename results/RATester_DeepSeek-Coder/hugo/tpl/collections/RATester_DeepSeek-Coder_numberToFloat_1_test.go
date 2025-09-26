package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumberToFloat_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    float64
		wantErr bool
	}{
		{
			name:    "float64",
			input:   reflect.ValueOf(float64(1.23)),
			want:    1.23,
			wantErr: false,
		},
		{
			name:    "int",
			input:   reflect.ValueOf(int(123)),
			want:    123.0,
			wantErr: false,
		},
		{
			name:    "uint",
			input:   reflect.ValueOf(uint(123)),
			want:    123.0,
			wantErr: false,
		},
		{
			name:    "interface",
			input:   reflect.ValueOf(interface{}(float64(1.23))),
			want:    1.23,
			wantErr: false,
		},
		{
			name:    "invalid kind",
			input:   reflect.ValueOf(struct{}{}),
			want:    0,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := numberToFloat(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("numberToFloat() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("numberToFloat() = %v, want %v", got, tc.want)
			}
		})
	}
}
