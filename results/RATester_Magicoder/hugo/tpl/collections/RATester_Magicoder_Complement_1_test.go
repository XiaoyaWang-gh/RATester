package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComplement_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   []any
		want    any
		wantErr bool
	}{
		{
			name:  "Test case 1",
			input: []any{[]any{1, 2, 3}, 1},
			want:  []any{2, 3},
		},
		{
			name:  "Test case 2",
			input: []any{[]any{1, 2, 3}, 4},
			want:  []any{1, 2, 3},
		},
		{
			name:    "Test case 3",
			input:   []any{[]any{1, 2, 3}},
			wantErr: true,
		},
		{
			name:    "Test case 4",
			input:   []any{1, 2, 3},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Complement(tt.input...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Complement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complement() = %v, want %v", got, tt.want)
			}
		})
	}
}
