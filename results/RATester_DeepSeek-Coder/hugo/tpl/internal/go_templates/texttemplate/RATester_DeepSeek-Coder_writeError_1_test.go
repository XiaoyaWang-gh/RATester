package template

import (
	"errors"
	"fmt"
	"testing"
)

func TestWriteError_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				err: errors.New("test error"),
			},
			want: "test error",
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

			s := &state{}
			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(writeError); ok {
						if err.Err.Error() != tt.want {
							t.Errorf("writeError() = %v, want %v", err.Err.Error(), tt.want)
						}
					} else {
						t.Errorf("writeError() = %v, want %v", r, tt.want)
					}
				} else {
					t.Errorf("writeError() did not panic")
				}
			}()
			s.writeError(tt.args.err)
		})
	}
}
