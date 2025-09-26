package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestBindWith_1(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		Field string `form:"field"`
	}

	tests := []struct {
		name    string
		obj     any
		binding binding.Binding
		wantErr bool
	}{
		{
			name:    "valid binding",
			obj:     &testStruct{},
			binding: binding.Form,
			wantErr: false,
		},
		{
			name:    "invalid binding",
			obj:     &testStruct{},
			binding: binding.Form,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				Request: &http.Request{},
			}
			if err := c.BindWith(tt.obj, tt.binding); (err != nil) != tt.wantErr {
				t.Errorf("Context.BindWith() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
