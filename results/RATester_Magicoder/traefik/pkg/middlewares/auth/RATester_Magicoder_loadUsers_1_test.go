package auth

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadUsers_1(t *testing.T) {
	type args struct {
		fileName    string
		appendUsers []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				fileName:    "test.txt",
				appendUsers: []string{"user1", "user2"},
			},
			want:    []string{"user1", "user2"},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				fileName:    "",
				appendUsers: []string{"user3", "user4"},
			},
			want:    []string{"user3", "user4"},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				fileName:    "non-existent.txt",
				appendUsers: []string{"user5", "user6"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := loadUsers(tt.args.fileName, tt.args.appendUsers)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
