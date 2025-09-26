package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSplicedStrList_1(t *testing.T) {
	type args struct {
		headerValue string
		dst         []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{
				headerValue: "test1,test2,test3",
				dst:         make([]string, 0),
			},
			want: []string{"test1", "test2", "test3"},
		},
		{
			name: "Test Case 2",
			args: args{
				headerValue: "test1, test2 , test3",
				dst:         make([]string, 0),
			},
			want: []string{"test1", "test2", "test3"},
		},
		{
			name: "Test Case 3",
			args: args{
				headerValue: "",
				dst:         make([]string, 0),
			},
			want: nil,
		},
		{
			name: "Test Case 4",
			args: args{
				headerValue: "test1",
				dst:         make([]string, 0),
			},
			want: []string{"test1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getSplicedStrList(tt.args.headerValue, tt.args.dst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSplicedStrList() = %v, want %v", got, tt.want)
			}
		})
	}
}
