package redactor

import (
	"fmt"
	"testing"
)

func TestDo_2(t *testing.T) {
	type args struct {
		baseConfig      interface{}
		tag             string
		redactByDefault bool
		indent          bool
	}
	tests := []struct {
		name    string
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

			got, err := do(tt.args.baseConfig, tt.args.tag, tt.args.redactByDefault, tt.args.indent)
			if (err != nil) != tt.wantErr {
				t.Errorf("do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("do() = %v, want %v", got, tt.want)
			}
		})
	}
}
