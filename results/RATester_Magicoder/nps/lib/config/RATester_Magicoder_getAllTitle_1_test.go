package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetAllTitle_1(t *testing.T) {
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
				content: "[Title1]\n[Title2]\n[Title1]",
			},
			wantArr: nil,
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				content: "[Title1]\n[Title2]\nTitle3",
			},
			wantArr: []string{"[Title1]", "[Title2]"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotArr, err := getAllTitle(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("getAllTitle() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
