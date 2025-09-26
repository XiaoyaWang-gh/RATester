package binding

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeJSON_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	tests := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid JSON",
			input: `{"field1": "value1", "field2": 2}`,
			want:  testStruct{Field1: "value1", Field2: 2},
		},
		{
			name:    "invalid JSON",
			input:   `{"field1": "value1", "field3": 3}`,
			wantErr: true,
		},
		{
			name:  "empty JSON",
			input: `{}`,
			want:  testStruct{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := strings.NewReader(tt.input)
			got := &testStruct{}
			err := decodeJSON(r, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("decodeJSON() got = %v, want %v", *got, tt.want)
			}
		})
	}
}
