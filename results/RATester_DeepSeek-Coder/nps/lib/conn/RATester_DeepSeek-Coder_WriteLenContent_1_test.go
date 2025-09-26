package conn

import (
	"fmt"
	"testing"
)

func TestWriteLenContent_1(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		s       *Conn
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

			if err := tt.s.WriteLenContent(tt.args.buf); (err != nil) != tt.wantErr {
				t.Errorf("Conn.WriteLenContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
