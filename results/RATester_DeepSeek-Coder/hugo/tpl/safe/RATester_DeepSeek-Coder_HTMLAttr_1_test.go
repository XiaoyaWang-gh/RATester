package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestHTMLAttr_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name    string
		s       any
		want    template.HTMLAttr
		wantErr bool
	}{
		{
			name:    "string",
			s:       "test",
			want:    "test",
			wantErr: false,
		},
		{
			name:    "int",
			s:       123,
			want:    "123",
			wantErr: false,
		},
		{
			name:    "float",
			s:       123.456,
			want:    "123.456",
			wantErr: false,
		},
		{
			name:    "nil",
			s:       nil,
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

			got, err := ns.HTMLAttr(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTMLAttr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTMLAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}
