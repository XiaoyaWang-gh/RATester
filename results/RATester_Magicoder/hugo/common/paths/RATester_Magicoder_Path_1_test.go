package paths

import (
	"fmt"
	"testing"
)

func TestPath_1(t *testing.T) {
	type fields struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		wantD  string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				s: "test",
			},
			wantD: "test",
		},
		{
			name: "Test Case 2",
			fields: fields{
				s: "test2",
			},
			wantD: "test2",
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
				s: tt.fields.s,
			}
			if gotD := p.Path(); gotD != tt.wantD {
				t.Errorf("Path() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}
