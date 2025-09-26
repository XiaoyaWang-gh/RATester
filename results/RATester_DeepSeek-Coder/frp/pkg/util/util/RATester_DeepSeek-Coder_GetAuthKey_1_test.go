package util

import (
	"fmt"
	"testing"
)

func TestGetAuthKey_1(t *testing.T) {
	type args struct {
		token     string
		timestamp int64
	}
	tests := []struct {
		name    string
		args    args
		wantKey string
	}{
		{
			name: "Test case 1",
			args: args{
				token:     "test_token",
				timestamp: 1638316800,
			},
			wantKey: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
		},
		{
			name: "Test case 2",
			args: args{
				token:     "another_test_token",
				timestamp: 1638403200,
			},
			wantKey: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotKey := GetAuthKey(tt.args.token, tt.args.timestamp); gotKey != tt.wantKey {
				t.Errorf("GetAuthKey() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
