package echo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUnwrap_4(t *testing.T) {
	type fields struct {
		Writer      http.ResponseWriter
		echo        *Echo
		beforeFuncs []func()
		afterFuncs  []func()
		Status      int
		Size        int64
		Committed   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   http.ResponseWriter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &Response{
				Writer:      tt.fields.Writer,
				echo:        tt.fields.echo,
				beforeFuncs: tt.fields.beforeFuncs,
				afterFuncs:  tt.fields.afterFuncs,
				Status:      tt.fields.Status,
				Size:        tt.fields.Size,
				Committed:   tt.fields.Committed,
			}
			if got := r.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
