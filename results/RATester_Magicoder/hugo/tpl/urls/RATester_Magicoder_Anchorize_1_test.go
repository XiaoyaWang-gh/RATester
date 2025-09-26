package urls

import (
	"fmt"
	"testing"
)

func TestAnchorize_1(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name    string
		ns      *Namespace
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			ns:   &Namespace{},
			args: args{
				s: "test",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Test case 2",
			ns:   &Namespace{},
			args: args{
				s: 123,
			},
			want:    "",
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

			got, err := tt.ns.Anchorize(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Anchorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.Anchorize() = %v, want %v", got, tt.want)
			}
		})
	}
}
