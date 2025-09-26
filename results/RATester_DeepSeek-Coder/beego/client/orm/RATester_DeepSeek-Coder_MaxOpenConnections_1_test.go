package orm

import (
	"fmt"
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

			got := MaxOpenConnections(tt.args.maxOpenConn)
			al := &alias{}
			got(al)
			if al.MaxOpenConns != tt.args.maxOpenConn {
				t.Errorf("MaxOpenConnections() = %v, want %v", al.MaxOpenConns, tt.args.maxOpenConn)
			}
		})
	}
}
