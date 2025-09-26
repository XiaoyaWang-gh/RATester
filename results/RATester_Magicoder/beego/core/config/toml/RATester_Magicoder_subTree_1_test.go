package toml

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestsubTree_1(t *testing.T) {
	type args struct {
		t    *toml.Tree
		path []string
	}
	tests := []struct {
		name    string
		args    args
		want    *toml.Tree
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				t:    &toml.Tree{},
				path: []string{"key1", "key2"},
			},
			want:    &toml.Tree{},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				t:    &toml.Tree{},
				path: []string{"key3", "key4"},
			},
			want:    nil,
			wantErr: true,
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

			got, err := subTree(tt.args.t, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("subTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
