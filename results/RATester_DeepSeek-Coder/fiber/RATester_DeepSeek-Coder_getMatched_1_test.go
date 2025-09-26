package fiber

import (
	"fmt"
	"testing"
)

func TestGetMatched_1(t *testing.T) {
	type fields struct {
		matched bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				matched: true,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				matched: false,
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

			c := &DefaultCtx{
				matched: tt.fields.matched,
			}
			if got := c.getMatched(); got != tt.want {
				t.Errorf("getMatched() = %v, want %v", got, tt.want)
			}
		})
	}
}
