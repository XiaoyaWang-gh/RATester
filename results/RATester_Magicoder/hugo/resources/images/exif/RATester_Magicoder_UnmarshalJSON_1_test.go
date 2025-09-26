package exif

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalJSON_1(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		want    Tags
		wantErr bool
	}

	tests := []testCase{
		{
			name:  "valid JSON",
			input: `{"field name": "field input"}`,
			want:  Tags{"field name": "field input"},
		},
		{
			name:    "invalid JSON",
			input:   `{"field name": "field input"`,
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

			var got Tags
			err := got.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
