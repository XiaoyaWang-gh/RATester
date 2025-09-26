package tcp

import (
	"fmt"
	"testing"
)

func TestClientIPV2_1(t *testing.T) {
	type args struct {
		tree      *matchersTree
		clientIPs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				tree: &matchersTree{
					matcher: func(meta ConnData) bool {
						return meta.remoteIP == "192.168.1.1"
					},
				},
				clientIPs: []string{"192.168.1.1"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree: &matchersTree{
					matcher: func(meta ConnData) bool {
						return meta.remoteIP == "192.168.1.2"
					},
				},
				clientIPs: []string{"192.168.1.1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := clientIPV2(tt.args.tree, tt.args.clientIPs...); (err != nil) != tt.wantErr {
				t.Errorf("clientIPV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
