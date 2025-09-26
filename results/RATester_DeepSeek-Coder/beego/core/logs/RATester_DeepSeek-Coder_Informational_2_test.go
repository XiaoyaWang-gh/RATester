package logs

import (
	"fmt"
	"testing"
)

func TestInformational_2(t *testing.T) {
	type args struct {
		f interface{}
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestInformational",
			args: args{
				f: "TestInformational",
				v: []interface{}{"TestInformational"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Informational(tt.args.f, tt.args.v...)
		})
	}
}
