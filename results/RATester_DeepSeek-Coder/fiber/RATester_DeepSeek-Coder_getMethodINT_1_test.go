package fiber

import (
	"fmt"
	"testing"
)

func TestGetMethodINT_1(t *testing.T) {
	type fields struct {
		methodINT int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test Case 1",
			fields: fields{
				methodINT: 1,
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			fields: fields{
				methodINT: 2,
			},
			want: 2,
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
				methodINT: tt.fields.methodINT,
			}
			if got := c.getMethodINT(); got != tt.want {
				t.Errorf("DefaultCtx.getMethodINT() = %v, want %v", got, tt.want)
			}
		})
	}
}
