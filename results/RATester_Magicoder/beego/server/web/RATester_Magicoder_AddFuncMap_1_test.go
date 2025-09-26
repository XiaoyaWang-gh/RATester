package web

import (
	"fmt"
	"testing"
)

func TestAddFuncMap_1(t *testing.T) {
	type args struct {
		key string
		fn  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestAddFuncMap_0",
			args: args{
				key: "testKey",
				fn:  func(a, b int) int { return a + b },
			},
			wantErr: false,
		},
		{
			name: "TestAddFuncMap_1",
			args: args{
				key: "testKey",
				fn:  func(a, b int) int { return a + b },
			},
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

			if err := AddFuncMap(tt.args.key, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("AddFuncMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
