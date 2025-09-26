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
			name:      "Test case 1",
			input:     "hello,world",
			delimiter: ",",
			want:      []string{"hello", "world"},
			wantErr:   false,
		},
		{
			name:      "Test case 2",
			input:     "hello world",
			delimiter: " ",
			want:      []string{"hello", "world"},
			wantErr:   false,
		},
		{
			name:      "Test case 3",
			input:     "hello,world",
			delimiter: ",",
			want:      []string{"hello", "world"},
			wantErr:   false,
		},
		{
			name:      "Test case 4",
			input:     "hello,world",
			delimiter: ",",
			want:      []string{"hello", "world"},
			wantErr:   false,
		},
		{
			name:      "Test case 5",
			input:     "hello,world",
			delimiter: ",",
			want:      []string{"hello", "world"},
			wantErr:   false,
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
