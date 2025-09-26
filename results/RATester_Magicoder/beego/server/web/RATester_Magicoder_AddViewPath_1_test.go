package web

import (
	"fmt"
	"testing"
)

func TestAddViewPath_1(t *testing.T) {
	type args struct {
		viewPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				viewPath: "test_path",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				viewPath: "test_path_2",
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := AddViewPath(tt.args.viewPath); (err != nil) != tt.wantErr {
				t.Errorf("AddViewPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
