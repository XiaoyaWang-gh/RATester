package related

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToKeyword_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		cfg  IndexConfig
		args args
		want Keyword
	}{
		{
			name: "lower case",
			cfg: IndexConfig{
				ToLower: true,
			},
			args: args{
				s: "FOO",
			},
			want: StringKeyword("foo"),
		},
		{
			name: "no lower case",
			cfg: IndexConfig{
				ToLower: false,
			},
			args: args{
				s: "FOO",
			},
			want: StringKeyword("FOO"),
		},
		{
			name: "fragments",
			cfg: IndexConfig{
				Type: TypeFragments,
			},
			args: args{
				s: "FOO",
			},
			want: FragmentKeyword("FOO"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.cfg.stringToKeyword(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}
