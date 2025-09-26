package pageparser

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestParseMain_1(t *testing.T) {
	type args struct {
		r   io.Reader
		cfg Config
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

			got, err := ParseMain(tt.args.r, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMain() = %v, want %v", got, tt.want)
			}
		})
	}
}
