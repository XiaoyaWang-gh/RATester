package fiber

import (
	"fmt"
	"testing"
)

func TestParamsMatch_1(t *testing.T) {
	type args struct {
		specParamStr headerParams
		offerParams  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				specParamStr: headerParams{"Content-Type": []byte("application/json")},
				offerParams:  "Content-Type=application/json",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				specParamStr: headerParams{"Content-Type": []byte("application/json"), "Accept": []byte("*/*")},
				offerParams:  "Content-Type=application/json; Accept=*/*",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				specParamStr: headerParams{"Content-Type": []byte("application/json")},
				offerParams:  "Content-Type=text/plain",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				specParamStr: headerParams{},
				offerParams:  "Content-Type=application/json",
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				specParamStr: headerParams{"Content-Type": []byte("application/json")},
				offerParams:  "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := paramsMatch(tt.args.specParamStr, tt.args.offerParams); got != tt.want {
				t.Errorf("paramsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
