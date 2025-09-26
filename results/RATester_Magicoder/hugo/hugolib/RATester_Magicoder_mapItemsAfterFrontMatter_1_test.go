package hugolib

import (
	"fmt"
	"testing"
)

func TestmapItemsAfterFrontMatter_1(t *testing.T) {
	type args struct {
		source []byte
		s      *shortcodeHandler
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

			if err := tt.rn.mapItemsAfterFrontMatter(tt.args.source, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("mapItemsAfterFrontMatter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
