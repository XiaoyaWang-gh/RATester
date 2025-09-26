package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDealMultiUser_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "testcase1",
			args: args{s: "a=1,b=2,c=3"},
			want: map[string]string{"a": "1", "b": "2", "c": "3"},
		},
		{
			name: "testcase2",
			args: args{s: "a=1,b=2,c=3,d=4"},
			want: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4"},
		},
		{
			name: "testcase3",
			args: args{s: "a=1,b=2,c=3,d=4,e=5"},
			want: map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := dealMultiUser(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealMultiUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
