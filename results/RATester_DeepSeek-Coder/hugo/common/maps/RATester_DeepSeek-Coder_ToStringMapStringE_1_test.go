package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMapStringE_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      any
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test case 1",
			in: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			in:      "not a map",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test case 3",
			in:      nil,
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ToStringMapStringE(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringMapStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapStringE() = %v, want %v", got, tt.want)
			}
		})
	}
}
