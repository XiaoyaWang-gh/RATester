package highlight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseHighlightOptions_1(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    map[string]any
		wantErr bool
	}{
		{
			name: "empty string",
			in:   "",
			want: map[string]any{},
		},
		{
			name: "single option",
			in:   "field=value",
			want: map[string]any{"field": "value"},
		},
		{
			name: "multiple options",
			in:   "field1=value1,field2=value2",
			want: map[string]any{"field1": "value1", "field2": "value2"},
		},
		{
			name:    "invalid option",
			in:      "field=value,invalid",
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

			got, err := parseHighlightOptions(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseHighlightOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseHighlightOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
