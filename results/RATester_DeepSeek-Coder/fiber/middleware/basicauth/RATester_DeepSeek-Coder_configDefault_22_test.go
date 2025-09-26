package basicauth

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigDefault_22(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test with no config",
			args: args{
				config: []Config{},
			},
			want: ConfigDefault,
		},
		{
			name: "Test with config",
			args: args{
				config: []Config{
					{
						Users: map[string]string{
							"user1": "pass1",
							"user2": "pass2",
						},
						Realm: "Test Realm",
					},
				},
			},
			want: Config{
				Next:         ConfigDefault.Next,
				Users:        map[string]string{"user1": "pass1", "user2": "pass2"},
				Realm:        "Test Realm",
				Authorizer:   ConfigDefault.Authorizer,
				Unauthorized: ConfigDefault.Unauthorized,
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

			got := configDefault(tt.args.config...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
