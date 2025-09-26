package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestinsertPageWithLock_1(t *testing.T) {
	type args struct {
		s string
		p *pageState
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
			gotU, gotN, gotB := m.insertPageWithLock(tt.args.s, tt.args.p)
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("insertPageWithLock() gotU = %v, want %v", gotU, tt.wantU)
			}
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("insertPageWithLock() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotB != tt.wantB {
				t.Errorf("insertPageWithLock() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
