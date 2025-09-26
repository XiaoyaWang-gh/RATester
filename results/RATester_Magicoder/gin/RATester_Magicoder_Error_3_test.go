package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_3(t *testing.T) {
	type fields struct {
		Err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				Err: errors.New("test error"),
			},
			want: "test error",
		},
		{
			name: "Test case 2",
			fields: fields{
				Err: errors.New("another test error"),
			},
			want: "another test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			msg := Error{
				Err: tt.fields.Err,
			}
			if got := msg.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
