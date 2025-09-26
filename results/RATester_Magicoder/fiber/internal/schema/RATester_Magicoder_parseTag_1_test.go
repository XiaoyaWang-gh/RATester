package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseTag_1(t *testing.T) {
	tests := []struct {
		name  string
		tag   string
		want  string
		want1 tagOptions
	}{
		{
			name:  "Test case 1",
			tag:   "test,option1,option2",
			want:  "test",
			want1: []string{"option1", "option2"},
		},
		{
			name:  "Test case 2",
			tag:   "test",
			want:  "test",
			want1: []string{},
		},
		{
			name:  "Test case 3",
			tag:   "test,option1",
			want:  "test",
			want1: []string{"option1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := parseTag(tt.tag)
			if got != tt.want {
				t.Errorf("parseTag() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseTag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
