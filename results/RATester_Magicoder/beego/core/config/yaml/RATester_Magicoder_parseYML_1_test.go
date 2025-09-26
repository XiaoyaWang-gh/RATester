package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseYML_1(t *testing.T) {
	tests := []struct {
		name    string
		buf     []byte
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test case 1",
			buf: []byte(`
				name: test
			`),
			want: map[string]interface{}{
				"name": "test",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			buf: []byte(`
				name: test
				age: 25
			`),
			want: map[string]interface{}{
				"name": "test",
				"age":  25,
			},
			wantErr: false,
		},
		{
			name:    "Test case 3",
			buf:     []byte(`invalid yaml`),
			want:    nil,
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
