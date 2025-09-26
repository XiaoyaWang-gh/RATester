package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSet_12(t *testing.T) {
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		e    *DateField
		args args
		want time.Time
	}{
		{
			name: "TestSet",
			e:    &DateField{},
			args: args{
				d: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.e.Set(tt.args.d)
			if !reflect.DeepEqual(time.Time(*tt.e), tt.want) {
				t.Errorf("Set() = %v, want %v", time.Time(*tt.e), tt.want)
			}
		})
	}
}
