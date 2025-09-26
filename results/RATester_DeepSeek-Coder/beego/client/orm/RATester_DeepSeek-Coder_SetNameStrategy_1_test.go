package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestSetNameStrategy_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test SetNameStrategy with valid strategy",
			args: args{s: "snake_case"},
			want: "snake_case",
		},
		{
			name: "Test SetNameStrategy with invalid strategy",
			args: args{s: "invalid_strategy"},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetNameStrategy(tt.args.s)
			if models.NameStrategy != tt.want {
				t.Errorf("SetNameStrategy() = %v, want %v", models.NameStrategy, tt.want)
			}
		})
	}
}
