package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdoIndex_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		item    any
		args    []any
		want    any
		wantErr bool
	}{
		{
			name: "Test case 1",
			item: map[string]any{"key": "value"},
			args: []any{"key"},
			want: "value",
		},
		{
			name: "Test case 2",
			item: []any{"item1", "item2", "item3"},
			args: []any{1},
			want: "item2",
		},
		{
			name:    "Test case 3",
			item:    nil,
			args:    []any{"key"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test case 4",
			item:    "string",
			args:    []any{1},
			want:    nil,
			wantErr: true,
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

			got, err := ns.doIndex(tt.item, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("doIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
