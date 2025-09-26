package goroutine

import (
	"fmt"
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestNewConnGroup_1(t *testing.T) {
	type args struct {
		dst io.ReadWriteCloser
		src io.ReadWriteCloser
		wg  *sync.WaitGroup
		n   *int64
	}
	tests := []struct {
		name string
		args args
		want connGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newConnGroup(tt.args.dst, tt.args.src, tt.args.wg, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newConnGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
