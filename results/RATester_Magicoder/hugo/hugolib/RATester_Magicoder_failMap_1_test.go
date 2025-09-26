package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/pageparser"
)

func TestfailMap_1(t *testing.T) {
	type args struct {
		source []byte
		err    error
		i      pageparser.Item
	}
	tests := []struct {
		name    string
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

			rn := &contentParseInfo{}
			if err := rn.failMap(tt.args.source, tt.args.err, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("failMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
