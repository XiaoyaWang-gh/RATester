package blockquotes

import (
	"fmt"
	"testing"
)

func TestAlertType_1(t *testing.T) {
	type fields struct {
		typ   string
		alert blockQuoteAlert
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				typ: "info",
				alert: blockQuoteAlert{
					typ:   "info",
					sign:  "+",
					title: "Information",
				},
			},
			want: "info",
		},
		{
			name: "Test case 2",
			fields: fields{
				typ: "warning",
				alert: blockQuoteAlert{
					typ:   "warning",
					sign:  "-",
					title: "Warning",
				},
			},
			want: "warning",
		},
		{
			name: "Test case 3",
			fields: fields{
				typ: "error",
				alert: blockQuoteAlert{
					typ:   "error",
					sign:  "x",
					title: "Error",
				},
			},
			want: "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &blockquoteContext{
				typ:   tt.fields.typ,
				alert: tt.fields.alert,
			}
			if got := c.AlertType(); got != tt.want {
				t.Errorf("blockquoteContext.AlertType() = %v, want %v", got, tt.want)
			}
		})
	}
}
