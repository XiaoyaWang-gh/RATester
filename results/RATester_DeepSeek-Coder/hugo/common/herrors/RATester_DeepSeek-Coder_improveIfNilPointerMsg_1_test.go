package herrors

import (
	"fmt"
	"testing"
)

func TestImproveIfNilPointerMsg_1(t *testing.T) {
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
				inErr: fmt.Errorf("nil pointer dereference: %s.%s", "receiver", "field"),
			},
			want: "– receiver is nil; wrap it in if or with: {{ with receiver }}{{ .field }}{{ end }}",
		},
		{
			name: "Test case 2",
			args: args{
				inErr: fmt.Errorf("nil pointer dereference: %s.%s", "anotherReceiver", "anotherField"),
			},
			want: "– anotherReceiver is nil; wrap it in if or with: {{ with anotherReceiver }}{{ .anotherField }}{{ end }}",
		},
		{
			name: "Test case 3",
			args: args{
				inErr: fmt.Errorf("some other error"),
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
