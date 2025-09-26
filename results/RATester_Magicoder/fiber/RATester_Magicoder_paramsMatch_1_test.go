package fiber

import (
	"fmt"
	"testing"
)

func TestparamsMatch_1(t *testing.T) {
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
				specParamStr: headerParams{"key1": []byte("value1")},
				offerParams:  "key1=value1",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				specParamStr: headerParams{"key1": []byte("value1")},
				offerParams:  "key1=value2",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				specParamStr: headerParams{"key1": []byte("value1"), "key2": []byte("value2")},
				offerParams:  "key1=value1&key2=value2",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				specParamStr: headerParams{"key1": []byte("value1"), "key2": []byte("value2")},
				offerParams:  "key1=value1&key3=value3",
			},
			want: false,
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
