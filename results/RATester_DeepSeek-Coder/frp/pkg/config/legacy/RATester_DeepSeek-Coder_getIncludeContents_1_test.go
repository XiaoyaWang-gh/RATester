package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetIncludeContents_1(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				paths: []string{"path1", "path2"},
			},
			want:    []byte("expected output"),
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				paths: []string{"path3", "path4"},
			},
			want:    []byte("expected output"),
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

			got, err := getIncludeContents(tt.args.paths)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIncludeContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIncludeContents() = %v, want %v", got, tt.want)
			}
		})
	}
}
