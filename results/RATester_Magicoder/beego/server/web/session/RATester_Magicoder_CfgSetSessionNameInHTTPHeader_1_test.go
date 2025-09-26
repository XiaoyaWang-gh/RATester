package session

import (
	"fmt"
	"testing"
)

func TestCfgSetSessionNameInHTTPHeader_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "Test case 1",
			args: args{
				name: "test_name",
			},
			want: func(config *ManagerConfig) {
				config.SessionNameInHTTPHeader = "test_name"
			},
		},
		{
			name: "Test case 2",
			args: args{
				name: "another_test_name",
			},
			want: func(config *ManagerConfig) {
				config.SessionNameInHTTPHeader = "another_test_name"
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

			got := CfgSetSessionNameInHTTPHeader(tt.args.name)
			config := &ManagerConfig{}
			got(config)
			if config.SessionNameInHTTPHeader != tt.args.name {
				t.Errorf("CfgSetSessionNameInHTTPHeader() = %v, want %v", config.SessionNameInHTTPHeader, tt.args.name)
			}
		})
	}
}
