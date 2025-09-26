package logger

import (
	"fmt"
	"testing"
)

func TestAppendInt_1(t *testing.T) {
	type args struct {
		output Buffer
		v      int
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

			got, err := appendInt(tt.args.output, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("appendInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("appendInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
