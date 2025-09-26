package time

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gohugoio/hugo/common/htime"
)

func TestNew_44(t *testing.T) {
	type args struct {
		timeFormatter htime.TimeFormatter
		location      *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Namespace
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

			if got := New(tt.args.timeFormatter, tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
