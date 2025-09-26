package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcaseInsensitiveLookup_1(t *testing.T) {
	type args struct {
		m reflect.Value
		k reflect.Value
	}
	tests := []struct {
		name  string
		args  args
		want  reflect.Value
		want1 bool
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

			got, got1 := caseInsensitiveLookup(tt.args.m, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caseInsensitiveLookup() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("caseInsensitiveLookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
