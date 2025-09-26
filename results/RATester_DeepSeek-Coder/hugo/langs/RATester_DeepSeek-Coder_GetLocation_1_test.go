package langs

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetLocation_1(t *testing.T) {
	type args struct {
		l *Language
	}
	tests := []struct {
		name string
		args args
		want *time.Location
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

			if got := GetLocation(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
