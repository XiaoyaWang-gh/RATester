package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxStmtCacheSize_1(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want DBOption
	}{
		{
			name: "Test case 1",
			args: args{v: 10},
			want: MaxStmtCacheSize(10),
		},
		{
			name: "Test case 2",
			args: args{v: 20},
			want: MaxStmtCacheSize(20),
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MaxStmtCacheSize(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxStmtCacheSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
