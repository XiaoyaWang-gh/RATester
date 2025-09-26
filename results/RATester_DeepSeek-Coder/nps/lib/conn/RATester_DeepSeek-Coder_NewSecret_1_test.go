package conn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSecret_1(t *testing.T) {
	type args struct {
		p    string
		conn *Conn
	}
	tests := []struct {
		name string
		args args
		want *Secret
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

			if got := NewSecret(tt.args.p, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
