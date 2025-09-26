package runtime

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdateServerStatus_1(t *testing.T) {
	type args struct {
		server string
		status string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test case 1",
			args: args{
				server: "server1",
				status: "status1",
			},
			want: map[string]string{
				"server1": "status1",
			},
		},
		{
			name: "Test case 2",
			args: args{
				server: "server2",
				status: "status2",
			},
			want: map[string]string{
				"server1": "status1",
				"server2": "status2",
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

			s := &ServiceInfo{
				serverStatus: make(map[string]string),
			}
			s.UpdateServerStatus(tt.args.server, tt.args.status)
			if !reflect.DeepEqual(s.serverStatus, tt.want) {
				t.Errorf("UpdateServerStatus() = %v, want %v", s.serverStatus, tt.want)
			}
		})
	}
}
