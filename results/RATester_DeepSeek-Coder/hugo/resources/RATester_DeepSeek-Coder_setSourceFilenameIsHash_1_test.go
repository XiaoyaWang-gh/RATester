package resources

import (
	"fmt"
	"testing"
)

func TestSetSourceFilenameIsHash_1(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		l    *genericResource
		args args
		want bool
	}{
		{
			name: "Test case 1",
			l:    &genericResource{sourceFilenameIsHash: false},
			args: args{b: true},
			want: true,
		},
		{
			name: "Test case 2",
			l:    &genericResource{sourceFilenameIsHash: true},
			args: args{b: false},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.l.setSourceFilenameIsHash(tt.args.b)
			if tt.l.sourceFilenameIsHash != tt.want {
				t.Errorf("setSourceFilenameIsHash() = %v, want %v", tt.l.sourceFilenameIsHash, tt.want)
			}
		})
	}
}
