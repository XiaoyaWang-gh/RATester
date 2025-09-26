package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewP2PServer_1(t *testing.T) {
	type args struct {
		p2pPort int
	}
	tests := []struct {
		name string
		args args
		want *P2PServer
	}{
		{
			name: "Test case 1",
			args: args{
				p2pPort: 8080,
			},
			want: &P2PServer{
				p2pPort: 8080,
				p2p:     make(map[string]*p2p),
			},
		},
		{
			name: "Test case 2",
			args: args{
				p2pPort: 8081,
			},
			want: &P2PServer{
				p2pPort: 8081,
				p2p:     make(map[string]*p2p),
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

			if got := NewP2PServer(tt.args.p2pPort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewP2PServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
