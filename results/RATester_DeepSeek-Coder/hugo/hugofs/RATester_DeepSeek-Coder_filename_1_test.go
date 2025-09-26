package hugofs

import (
	"fmt"
	"testing"
)

func TestFilename_1(t *testing.T) {
	type test struct {
		name    string
		r       RootMapping
		arg     string
		want    string
		wantErr bool
	}

	tests := []test{
		{
			name: "Test with empty name",
			r: RootMapping{
				From: "/from",
				To:   "/to",
			},
			arg:     "",
			want:    "/to",
			wantErr: false,
		},
		{
			name: "Test with non-empty name",
			r: RootMapping{
				From: "/from",
				To:   "/to",
			},
			arg:     "/from/file",
			want:    "/to/file",
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

			got := tt.r.filename(tt.arg)
			if got != tt.want {
				t.Errorf("filename() = %v, want %v", got, tt.want)
			}
		})
	}
}
