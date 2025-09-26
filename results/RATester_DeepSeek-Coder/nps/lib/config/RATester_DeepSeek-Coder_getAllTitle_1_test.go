package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAllTitle_1(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantArr []string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				content: "[Title1]\n[Title2]\n[Title3]",
			},
			wantArr: []string{"[Title1]", "[Title2]", "[Title3]"},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				content: "[Title1]\n[Title1]\n[Title3]",
			},
			wantArr: nil,
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				content: "[Title1]\r\n[Title2]\r\n[Title3]",
			},
			wantArr: []string{"[Title1]", "[Title2]", "[Title3]"},
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				content: "[Title1]\r\n[Title1]\r\n[Title3]",
			},
			wantArr: nil,
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

			gotArr, gotErr := getAllTitle(tt.args.content)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("getAllTitle() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("getAllTitle() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
