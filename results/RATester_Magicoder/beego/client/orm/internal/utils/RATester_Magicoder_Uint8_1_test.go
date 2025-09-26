package utils

import (
	"fmt"
	"testing"
)

func TestUint8_1(t *testing.T) {
	type args struct {
		s StrTo
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
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

			got, err := tt.args.s.Uint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
