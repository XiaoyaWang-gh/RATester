package context

import (
	"fmt"
	"testing"
)

func TestSetStatus_1(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestSetStatus",
			args: args{status: 200},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			output := &BeegoOutput{}
			output.SetStatus(tt.args.status)
			if output.Status != tt.want {
				t.Errorf("SetStatus() = %v, want %v", output.Status, tt.want)
			}
		})
	}
}
