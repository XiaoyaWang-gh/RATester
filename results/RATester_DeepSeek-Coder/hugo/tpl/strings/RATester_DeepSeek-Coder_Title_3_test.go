package strings

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestTitle_3(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestTitle_1",
			args: args{
				s: "test",
			},
			want:    "Test",
			wantErr: false,
		},
		{
			name: "TestTitle_2",
			args: args{
				s: 123,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "TestTitle_3",
			args: args{
				s: nil,
			},
			want:    "",
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

			ns := &Namespace{
				deps: &deps.Deps{},
			}
			got, err := ns.Title(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Title() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}
