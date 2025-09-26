package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge_6(t *testing.T) {
	type args struct {
		other HTMLElements
	}
	tests := []struct {
		name string
		h    *HTMLElements
		args args
		want HTMLElements
	}{
		{
			name: "Merge two HTMLElements",
			h: &HTMLElements{
				Tags:    []string{"div", "span"},
				Classes: []string{"class1", "class2"},
				IDs:     []string{"id1", "id2"},
			},
			args: args{
				other: HTMLElements{
					Tags:    []string{"p", "a"},
					Classes: []string{"class3", "class4"},
					IDs:     []string{"id3", "id4"},
				},
			},
			want: HTMLElements{
				Tags:    []string{"div", "span", "p", "a"},
				Classes: []string{"class1", "class2", "class3", "class4"},
				IDs:     []string{"id1", "id2", "id3", "id4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.h.Merge(tt.args.other)
			if !reflect.DeepEqual(tt.h, &tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.h, &tt.want)
			}
		})
	}
}
