package highlight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPreWrapper_1(t *testing.T) {
	type args struct {
		language     string
		writeCounter *byteCountFlexiWriter
	}
	tests := []struct {
		name string
		args args
		want *preWrapper
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

			if got := getPreWrapper(tt.args.language, tt.args.writeCounter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPreWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
