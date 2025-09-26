package web

import (
	"fmt"
	"testing"
)

func TestGt_1(t *testing.T) {
	type test struct {
		arg1, arg2 interface{}
		want       bool
		wantErr    bool
	}

	tests := []test{
		{arg1: 1, arg2: 2, want: false, wantErr: false},
		{arg1: 2, arg2: 1, want: true, wantErr: false},
		{arg1: 1, arg2: 1, want: false, wantErr: false},
		{arg1: "a", arg2: "b", want: false, wantErr: false},
		{arg1: "b", arg2: "a", want: true, wantErr: false},
		{arg1: "a", arg2: "a", want: false, wantErr: false},
		{arg1: 1, arg2: "a", want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("gt(%v, %v)", tt.arg1, tt.arg2), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := gt(tt.arg1, tt.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("gt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gt() = %v, want %v", got, tt.want)
			}
		})
	}
}
