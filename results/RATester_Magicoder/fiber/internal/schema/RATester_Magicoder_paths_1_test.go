package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Testpaths_1(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		f    *fieldInfo
		args args
		want []string
	}{
		{
			name: "Test case 1",
			f: &fieldInfo{
				name:             "field name",
				alias:            "field alias",
				canonicalAlias:   "field canonicalAlias",
				unmarshalerInfo:  unmarshaler{},
				isSliceOfStructs: false,
			},
			args: args{
				prefix: "prefix",
			},
			want: []string{"prefix + field alias"},
		},
		{
			name: "Test case 2",
			f: &fieldInfo{
				name:             "field name",
				alias:            "field alias",
				canonicalAlias:   "field canonicalAlias",
				unmarshalerInfo:  unmarshaler{},
				isSliceOfStructs: true,
			},
			args: args{
				prefix: "prefix",
			},
			want: []string{"prefix + field alias", "prefix + field canonicalAlias"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.paths(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldInfo.paths() = %v, want %v", got, tt.want)
			}
		})
	}
}
