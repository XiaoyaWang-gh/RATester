package strings

import (
	"fmt"
	"testing"
)

func TestRepeat_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		n       any
		s       any
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			n:       3,
			s:       "hello",
			want:    "hellohellohello",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			n:       -1,
			s:       "hello",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Test case 3",
			n:       0,
			s:       "hello",
			want:    "",
			wantErr: false,
		},
		{
			name:    "Test case 4",
			n:       3,
			s:       123,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Test case 5",
			n:       "3",
			s:       "hello",
			want:    "hellohellohello",
			wantErr: false,
		},
		{
			name:    "Test case 6",
			n:       "3",
			s:       123,
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

			got, err := ns.Repeat(tt.n, tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Repeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
