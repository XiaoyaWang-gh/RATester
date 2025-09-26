package paths

import (
	"fmt"
	"testing"
)

func TestComponent_1(t *testing.T) {
	type fields struct {
		component string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				component: "test",
			},
			want: "test",
		},
		{
			name: "Test Case 2",
			fields: fields{
				component: "test2",
			},
			want: "test2",
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
				component: tt.fields.component,
			}
			if got := p.Component(); got != tt.want {
				t.Errorf("Path.Component() = %v, want %v", got, tt.want)
			}
		})
	}
}
