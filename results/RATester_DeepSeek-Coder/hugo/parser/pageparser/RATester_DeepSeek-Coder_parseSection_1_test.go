package pageparser

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestParseSection_1(t *testing.T) {
	type args struct {
		r     io.Reader
		cfg   Config
		start stateFunc
	}
	tests := []struct {
		name    string
		args    args
		want    Result
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

			got, err := parseSection(tt.args.r, tt.args.cfg, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
