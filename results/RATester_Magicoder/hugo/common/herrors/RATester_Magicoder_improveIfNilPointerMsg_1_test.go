package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestimproveIfNilPointerMsg_1(t *testing.T) {
	type args struct {
		inErr error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				inErr: errors.New("nil pointer dereference: field name"),
			},
			want: "– field name is nil; wrap it in if or with: {{ with receiver }}{{ .field }}{{ end }}",
		},
		{
			name: "Test case 2",
			args: args{
				inErr: errors.New("some other error"),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := improveIfNilPointerMsg(tt.args.inErr); got != tt.want {
				t.Errorf("improveIfNilPointerMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
