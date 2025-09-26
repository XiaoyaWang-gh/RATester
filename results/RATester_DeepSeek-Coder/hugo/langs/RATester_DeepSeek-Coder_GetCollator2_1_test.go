package langs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCollator2_1(t *testing.T) {
	type args struct {
		l *Language
	}
	tests := []struct {
		name string
		args args
		want *Collator
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

			if got := GetCollator2(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollator2() = %v, want %v", got, tt.want)
			}
		})
	}
}
