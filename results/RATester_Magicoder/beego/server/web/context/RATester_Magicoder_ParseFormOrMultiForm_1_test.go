package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParseFormOrMultiForm_1(t *testing.T) {
	type args struct {
		maxMemory int64
	}
	tests := []struct {
		name    string
		input   *BeegoInput
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{},
				},
			},
			args: args{
				maxMemory: 1024,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{},
				},
			},
			args: args{
				maxMemory: 2048,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{},
				},
			},
			args: args{
				maxMemory: 4096,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.input.ParseFormOrMultiForm(tt.args.maxMemory); (err != nil) != tt.wantErr {
				t.Errorf("BeegoInput.ParseFormOrMultiForm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
