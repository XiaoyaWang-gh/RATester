package pageparser

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestParseFrontMatterAndContent_1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    ContentFrontMatter
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

			got, err := ParseFrontMatterAndContent(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFrontMatterAndContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFrontMatterAndContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
