package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxOpenConnections_1(t *testing.T) {
	type args struct {
		maxOpenConn int
	}
	tests := []struct {
		name string
		args args
		want DBOption
	}{
		{
			name: "Test case 1",
			args: args{
				maxOpenConn: 10,
			},
			want: func(al *alias) {
				al.SetMaxOpenConns(10)
			},
		},
		{
			name: "Test case 2",
			args: args{
				maxOpenConn: 20,
			},
			want: func(al *alias) {
				al.SetMaxOpenConns(20)
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

			if got := MaxOpenConnections(tt.args.maxOpenConn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxOpenConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}
