package tcp

import (
	"fmt"
	"testing"
)

func TestClientIP_1(t *testing.T) {
	type args struct {
		tree     *matchersTree
		clientIP []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			if err := clientIP(tt.args.tree, tt.args.clientIP...); (err != nil) != tt.wantErr {
				t.Errorf("clientIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
