package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNext_1(t *testing.T) {
	type fields struct {
		Major      int
		Minor      int
		PatchLevel int
		Suffix     string
	}
	tests := []struct {
		name   string
		fields fields
		want   Version
	}{
		{
			name: "TestNext_1",
			fields: fields{
				Major: 1,
				Minor: 2,
			},
			want: Version{
				Major: 1,
				Minor: 3,
			},
		},
		{
			name: "TestNext_2",
			fields: fields{
				Major: 2,
				Minor: 0,
			},
			want: Version{
				Major: 2,
				Minor: 1,
			},
		},
		{
			name: "TestNext_3",
			fields: fields{
				Major: 0,
				Minor: 0,
			},
			want: Version{
				Major: 0,
				Minor: 1,
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

			v := Version{
				Major:      tt.fields.Major,
				Minor:      tt.fields.Minor,
				PatchLevel: tt.fields.PatchLevel,
				Suffix:     tt.fields.Suffix,
			}
			if got := v.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Version.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
