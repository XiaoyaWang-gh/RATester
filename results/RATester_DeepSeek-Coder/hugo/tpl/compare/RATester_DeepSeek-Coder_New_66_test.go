package compare

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew_66(t *testing.T) {
	type args struct {
		loc             *time.Location
		caseInsensitive bool
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

			if got := New(tt.args.loc, tt.args.caseInsensitive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
