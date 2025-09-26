package hugolib

import (
	"fmt"
	"testing"
)

func TestmapFrontMatter_1(t *testing.T) {
	type args struct {
		source []byte
	}
	tests := []struct {
		name    string
		rn      *contentParseInfo
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

			if err := tt.rn.mapFrontMatter(tt.args.source); (err != nil) != tt.wantErr {
				t.Errorf("mapFrontMatter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
