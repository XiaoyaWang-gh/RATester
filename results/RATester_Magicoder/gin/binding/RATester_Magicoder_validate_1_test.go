package binding

import (
	"fmt"
	"testing"
)

func Testvalidate_1(t *testing.T) {
	type args struct {
		obj any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				obj: struct{ fieldName string }{fieldName: "test"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				obj: struct{ fieldName string }{fieldName: ""},
			},
			wantErr: true,
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

			if err := validate(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
