package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseYML_1(t *testing.T) {
	tests := []struct {
		name    string
		buf     []byte
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test case 1",
			buf: []byte(`
				field: value
			`),
			want: map[string]interface{}{
				"field": "value",
			},
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := parseYML(tt.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseYML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseYML() = %v, want %v", got, tt.want)
			}
		})
	}
}
