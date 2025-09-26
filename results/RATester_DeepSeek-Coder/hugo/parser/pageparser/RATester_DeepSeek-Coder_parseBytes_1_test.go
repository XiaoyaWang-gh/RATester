package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseBytes_1(t *testing.T) {
	type args struct {
		b     []byte
		cfg   Config
		start stateFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *pageLexer
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

			got, err := parseBytes(tt.args.b, tt.args.cfg, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
