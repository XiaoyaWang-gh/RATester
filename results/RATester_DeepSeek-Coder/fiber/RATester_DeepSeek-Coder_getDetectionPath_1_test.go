package fiber

import (
	"fmt"
	"testing"
)

func TestGetDetectionPath_1(t *testing.T) {
	type fields struct {
		detectionPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				detectionPath: "test",
			},
			want: "test",
		},
		{
			name: "Test Case 2",
			fields: fields{
				detectionPath: "test2",
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

			c := &DefaultCtx{
				detectionPath: tt.fields.detectionPath,
			}
			if got := c.getDetectionPath(); got != tt.want {
				t.Errorf("getDetectionPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
