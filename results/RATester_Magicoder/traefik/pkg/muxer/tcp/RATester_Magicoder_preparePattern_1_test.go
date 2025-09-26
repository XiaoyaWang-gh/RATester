package tcp

import (
	"fmt"
	"testing"
)

func TestPreparePattern_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "test case 1",
			input:   "{field name}:{field input}",
			want:    "(?i)^(?P<0>[^.]+){field name}(?P<1>[^.]+){field input}$",
			wantErr: false,
		},
		{
			name:    "test case 2",
			input:   "{field name}:{field input}:{field want}",
			want:    "(?i)^(?P<0>[^.]+){field name}(?P<1>[^.]+){field input}(?P<2>[^.]+){field want}$",
			wantErr: false,
		},
		{
			name:    "test case 3",
			input:   "{field name}:{field input}:{field want}:{field wantErr}",
			want:    "(?i)^(?P<0>[^.]+){field name}(?P<1>[^.]+){field input}(?P<2>[^.]+){field want}(?P<3>[^.]+){field wantErr}$",
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

			got, err := preparePattern(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("preparePattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("preparePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
