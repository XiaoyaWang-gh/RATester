package conn

import (
	"fmt"
	"testing"
)

func TestReadLen_1(t *testing.T) {
	type args struct {
		cLen int
		buf  []byte
	}
	tests := []struct {
		name    string
		s       *Conn
		args    args
		want    int
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

			got, err := tt.s.ReadLen(tt.args.cLen, tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.ReadLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Conn.ReadLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
