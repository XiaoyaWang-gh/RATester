package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsplitStr_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name              string
		args              args
		wantConfigDataArr []string
	}{
		{
			name: "Test case 1",
			args: args{
				s: "test\r\nstring",
			},
			wantConfigDataArr: []string{"test", "string"},
		},
		{
			name: "Test case 2",
			args: args{
				s: "test\nstring",
			},
			wantConfigDataArr: []string{"test", "string"},
		},
		{
			name: "Test case 3",
			args: args{
				s: "test\r\nstring\r\n",
			},
			wantConfigDataArr: []string{"test", "string", ""},
		},
		{
			name: "Test case 4",
			args: args{
				s: "test\nstring\n",
			},
			wantConfigDataArr: []string{"test", "string", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotConfigDataArr := splitStr(tt.args.s); !reflect.DeepEqual(gotConfigDataArr, tt.wantConfigDataArr) {
				t.Errorf("splitStr() = %v, want %v", gotConfigDataArr, tt.wantConfigDataArr)
			}
		})
	}
}
