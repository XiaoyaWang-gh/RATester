package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromContent_1(t *testing.T) {
	type args struct {
		types          Types
		extensionHints []string
		content        []byte
	}
	tests := []struct {
		name string
		args args
		want Type
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

			if got := FromContent(tt.args.types, tt.args.extensionHints, tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
