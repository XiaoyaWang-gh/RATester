package warpc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllDispatchers_1(t *testing.T) {
	type args struct {
		katexOpts Options
	}
	tests := []struct {
		name string
		args args
		want *Dispatchers
	}{
		{
			name: "Test Case 1",
			args: args{
				katexOpts: Options{
					Runtime: Binary{Name: "javy_quickjs_provider_v2", Data: quickjsWasm},
					Main:    Binary{Name: "renderkatex", Data: katexWasm},
					Infof: func(format string, v ...any) {

					},
				},
			},
			want: &Dispatchers{
				katex: &lazyDispatcher[KatexInput, KatexOutput]{opts: Options{
					Runtime: Binary{Name: "javy_quickjs_provider_v2", Data: quickjsWasm},
					Main:    Binary{Name: "renderkatex", Data: katexWasm},
					Infof: func(format string, v ...any) {

					},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AllDispatchers(tt.args.katexOpts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllDispatchers() = %v, want %v", got, tt.want)
			}
		})
	}
}
