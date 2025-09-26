package echo

import (
	"fmt"
	"testing"
)

func TestError_4(t *testing.T) {
	type fields struct {
		Field     string
		HTTPError *HTTPError
		Values    []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Error",
			fields: fields{
				Field: "testField",
				HTTPError: &HTTPError{
					Message: "testMessage",
					Code:    400,
				},
				Values: []string{"value1", "value2"},
			},
			want: "testMessage, field=testField",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			be := &BindingError{
				Field:     tt.fields.Field,
				HTTPError: tt.fields.HTTPError,
				Values:    tt.fields.Values,
			}
			if got := be.Error(); got != tt.want {
				t.Errorf("BindingError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
