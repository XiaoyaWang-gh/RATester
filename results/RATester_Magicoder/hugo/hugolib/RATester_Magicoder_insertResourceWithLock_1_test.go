package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestinsertResourceWithLock_1(t *testing.T) {
	type args struct {
		s string
		r contentNodeI
	}
	tests := []struct {
		name  string
		args  args
		wantU contentNodeI
		wantN contentNodeI
		wantB bool
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

			m := &pageMap{}
			gotU, gotN, gotB := m.insertResourceWithLock(tt.args.s, tt.args.r)
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("insertResourceWithLock() gotU = %v, want %v", gotU, tt.wantU)
			}
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("insertResourceWithLock() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotB != tt.wantB {
				t.Errorf("insertResourceWithLock() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
