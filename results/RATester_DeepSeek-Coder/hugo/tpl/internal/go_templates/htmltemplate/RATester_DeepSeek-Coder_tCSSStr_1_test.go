package template

import (
	"fmt"
	"testing"
)

func TestTCSSStr_1(t *testing.T) {
	tests := []struct {
		name    string
		c       context
		s       []byte
		wantC   context
		wantN   int
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

			gotC, gotN := tCSSStr(tt.c, tt.s)
			if (gotC.state != tt.wantC.state) || (gotN != tt.wantN) {
				t.Errorf("tCSSStr() = %v, %v, want %v, %v", gotC, gotN, tt.wantC, tt.wantN)
			}
			if (gotC.err != nil) != tt.wantErr {
				t.Errorf("tCSSStr() error = %v, wantErr %v", gotC.err, tt.wantErr)
			}
		})
	}
}
