package session

import (
	"fmt"
	"testing"
)

func TestdecodeSessionData_1(t *testing.T) {
	type args struct {
		rawData []byte
	}
	tests := []struct {
		name    string
		s       *Session
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

			if err := tt.s.decodeSessionData(tt.args.rawData); (err != nil) != tt.wantErr {
				t.Errorf("Session.decodeSessionData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
