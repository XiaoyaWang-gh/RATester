package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitStr_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "testcase1",
			args: args{
				s: "testcase1",
			},
			want: []string{"testcase1"},
		},
		{
			name: "testcase2",
			args: args{
				s: "testcase2\ntestcase2",
			},
			want: []string{"testcase2", "testcase2"},
		},
		{
			name: "testcase3",
			args: args{
				s: "testcase3\r\ntestcase3",
			},
			want: []string{"testcase3", "testcase3"},
		},
		{
			name: "testcase4",
			args: args{
				s: "testcase4\ntestcase4\ntestcase4",
			},
			want: []string{"testcase4", "testcase4", "testcase4"},
		},
		{
			name: "testcase5",
			args: args{
				s: "testcase5\r\ntestcase5\r\ntestcase5",
			},
			want: []string{"testcase5", "testcase5", "testcase5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := splitStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
