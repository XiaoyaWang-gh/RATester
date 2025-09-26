package mock

import (
	"fmt"
	"testing"
)

func TestExist_2(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		d    *DoNothingQueryM2Mer
		args args
		want bool
	}{
		{
			name: "TestExist",
			d:    &DoNothingQueryM2Mer{},
			args: args{
				i: "test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.Exist(tt.args.i); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
