package langs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/htime"
)

func TestGetTimeFormatter_1(t *testing.T) {
	type args struct {
		l *Language
	}
	tests := []struct {
		name string
		args args
		want htime.TimeFormatter
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

			if got := GetTimeFormatter(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimeFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}
