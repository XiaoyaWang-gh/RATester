package urls

import (
	"fmt"
	"testing"
)

func TestRelRef_1(t *testing.T) {
	type args struct {
		p    any
		args any
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

			got, err := tt.ns.RelRef(tt.args.p, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.RelRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.RelRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
