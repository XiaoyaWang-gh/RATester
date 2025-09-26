package metadecoders

import (
	"errors"
	"fmt"
	"testing"
)

func TesttoFileError_1(t *testing.T) {
	type args struct {
		f    Format
		data []byte
		err  error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				f:    "testFormat",
				data: []byte("testData"),
				err:  errors.New("testError"),
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				f:    "anotherFormat",
				data: []byte("anotherData"),
				err:  nil,
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := toFileError(tt.args.f, tt.args.data, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("toFileError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
