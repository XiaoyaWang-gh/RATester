package hreflect

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestAsTime_2(t *testing.T) {
	type args struct {
		v   reflect.Value
		loc *time.Location
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
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

			got, got1 := AsTime(tt.args.v, tt.args.loc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AsTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
