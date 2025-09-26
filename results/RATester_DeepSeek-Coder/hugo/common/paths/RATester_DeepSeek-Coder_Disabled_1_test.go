package paths

import (
	"fmt"
	"testing"
)

func TestDisabled_1(t *testing.T) {
	type fields struct {
		disabled bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				disabled: true,
			},
			want: true,
		},
		{
			name: "Test case 2",
			fields: fields{
				disabled: false,
			},
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

			p := &Path{
				disabled: tt.fields.disabled,
			}
			if got := p.Disabled(); got != tt.want {
				t.Errorf("Path.Disabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
