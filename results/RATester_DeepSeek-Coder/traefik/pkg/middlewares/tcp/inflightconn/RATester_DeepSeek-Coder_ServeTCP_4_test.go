package inflightconn

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestServeTCP_4(t *testing.T) {
	type args struct {
		conn tcp.WriteCloser
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &inFlightConn{}
			i.ServeTCP(tt.args.conn)
		})
	}
}
