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
		name string
		args args
		want time.Time
	}{
		{
			name: "case 1",
			args: args{
				v:   reflect.ValueOf(time.Now()),
				loc: time.Local,
			},
			want: time.Now(),
		},
		{
			name: "case 2",
			args: args{
				v:   reflect.ValueOf(time.Now()),
				loc: time.UTC,
			},
			want: time.Now().UTC(),
		},
		{
			name: "case 3",
			args: args{
				v:   reflect.ValueOf(time.Now()),
				loc: nil,
			},
			want: time.Now(),
		},
		{
			name: "case 4",
			args: args{
				v:   reflect.ValueOf(time.Now()),
				loc: time.UTC,
			},
			want: time.Now().UTC(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got, _ := AsTime(tt.args.v, tt.args.loc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
