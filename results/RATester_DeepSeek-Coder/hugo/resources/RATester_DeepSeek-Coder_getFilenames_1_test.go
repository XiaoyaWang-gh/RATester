package resources

import (
	"fmt"
	"testing"
)

func TestGetFilenames_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "Test Case 1",
			args:  args{key: "testKey"},
			want:  "testKey.json",
			want1: "testKey.content",
		},
		{
			name:  "Test Case 2",
			args:  args{key: "anotherKey"},
			want:  "anotherKey.json",
			want1: "anotherKey.content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ResourceCache{}
			got, got1 := c.getFilenames(tt.args.key)
			if got != tt.want {
				t.Errorf("getFilenames() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getFilenames() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
