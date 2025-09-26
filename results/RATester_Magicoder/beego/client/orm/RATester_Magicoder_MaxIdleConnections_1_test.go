package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxIdleConnections_1(t *testing.T) {
	type args struct {
		maxIdleConn int
	}
	tests := []struct {
		name string
		args args
		want DBOption
	}{
		{
			name: "Test case 1",
			args: args{
				maxIdleConn: 10,
			},
			want: func(al *alias) {
				al.SetMaxIdleConns(10)
			},
		},
		{
			name: "Test case 2",
			args: args{
				maxIdleConn: 20,
			},
			want: func(al *alias) {
				al.SetMaxIdleConns(20)
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

			if got := MaxIdleConnections(tt.args.maxIdleConn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxIdleConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}
