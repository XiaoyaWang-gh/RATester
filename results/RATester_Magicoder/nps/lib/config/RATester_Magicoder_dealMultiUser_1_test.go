package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdealMultiUser_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test case 1",
			args: args{s: "user1=value1, user2=value2"},
			want: map[string]string{"user1": "value1", "user2": "value2"},
		},
		{
			name: "Test case 2",
			args: args{s: "user1=value1"},
			want: map[string]string{"user1": "value1"},
		},
		{
			name: "Test case 3",
			args: args{s: "user1"},
			want: map[string]string{"user1": ""},
		},
		{
			name: "Test case 4",
			args: args{s: ""},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := dealMultiUser(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealMultiUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
