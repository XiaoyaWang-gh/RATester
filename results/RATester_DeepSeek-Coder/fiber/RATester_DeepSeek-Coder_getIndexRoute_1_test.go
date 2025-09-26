package fiber

import (
	"fmt"
	"testing"
)

func TestGetIndexRoute_1(t *testing.T) {
	type fields struct {
		indexRoute int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test Case 1",
			fields: fields{
				indexRoute: 1,
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			fields: fields{
				indexRoute: 2,
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
				indexRoute: tt.fields.indexRoute,
			}
			if got := c.getIndexRoute(); got != tt.want {
				t.Errorf("DefaultCtx.getIndexRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
