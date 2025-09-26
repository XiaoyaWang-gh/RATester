package conn

import (
	"fmt"
	"testing"
	"time"
)

func TestSetReadDeadline_1(t *testing.T) {
	type args struct {
		deadline time.Time
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

			if err := tt.s.SetReadDeadline(tt.args.deadline); (err != nil) != tt.wantErr {
				t.Errorf("Conn.SetReadDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
