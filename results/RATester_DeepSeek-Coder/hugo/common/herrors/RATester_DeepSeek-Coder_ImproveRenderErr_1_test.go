package herrors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestImproveRenderErr_1(t *testing.T) {
	type args struct {
		inErr error
	}
	tests := []struct {
		name       string
		args       args
		wantOutErr error
	}{
		{
			name: "Test case 1",
			args: args{
				inErr: errors.New("test error"),
			},
			wantOutErr: errors.New("test error"),
		},
		{
			name: "Test case 2",
			args: args{
				inErr: &errMessage{msg: "test message", err: errors.New("test error")},
			},
			wantOutErr: &errMessage{msg: "test message", err: errors.New("test error")},
		},
		{
			name: "Test case 3",
			args: args{
				inErr: errors.New("test error: deferred"),
			},
			wantOutErr: &errMessage{msg: "executing test error", err: errors.New("test error")},
		},
		{
			name: "Test case 4",
			args: args{
				inErr: &errMessage{msg: "test message: deferred", err: errors.New("test error")},
			},
			wantOutErr: &errMessage{msg: "test message: executing test error", err: &errMessage{msg: "test message", err: errors.New("test error")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotOutErr := ImproveRenderErr(tt.args.inErr); !reflect.DeepEqual(gotOutErr, tt.wantOutErr) {
				t.Errorf("ImproveRenderErr() = %v, want %v", gotOutErr, tt.wantOutErr)
			}
		})
	}
}
