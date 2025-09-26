package conn

import (
	"fmt"
	"testing"
)

func TestSendInfo_1(t *testing.T) {
	type args struct {
		t    interface{}
		flag string
	}
	tests := []struct {
		name    string
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

			s := &Conn{}
			got, err := s.SendInfo(tt.args.t, tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
