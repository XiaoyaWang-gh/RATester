package os

import (
	"fmt"
	"testing"
)

func TestGetenv_1(t *testing.T) {
	type args struct {
		key any
	}
	tests := []struct {
		name    string
		ns      *Namespace
		args    args
		want    string
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

			got, err := tt.ns.Getenv(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Getenv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.Getenv() = %v, want %v", got, tt.want)
			}
		})
	}
}
