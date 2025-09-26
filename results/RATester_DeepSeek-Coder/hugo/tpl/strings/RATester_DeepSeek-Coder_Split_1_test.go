package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name      string
		input     any
		delimiter string
		want      []string
		wantErr   bool
	}{
		{
			name:      "Test with string",
			input:     "hello,world",
			delimiter: ",",
			want:      []string{"hello", "world"},
			wantErr:   false,
		},
		{
			name:      "Test with int",
			input:     123,
			delimiter: ",",
			want:      []string{},
			wantErr:   true,
		},
		{
			name:      "Test with float",
			input:     123.456,
			delimiter: ",",
			want:      []string{},
			wantErr:   true,
		},
		{
			name:      "Test with bool",
			input:     true,
			delimiter: ",",
			want:      []string{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Split(tt.input, tt.delimiter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
