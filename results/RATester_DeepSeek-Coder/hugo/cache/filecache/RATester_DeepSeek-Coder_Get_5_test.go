package filecache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_5(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 bool
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

			h := &httpCache{}
			got, got1 := h.Get(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("httpCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
