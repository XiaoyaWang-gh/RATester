package conn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetShortContent_1(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name    string
		s       *Conn
		args    args
		wantB   []byte
		wantErr bool
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

			gotB, err := tt.s.GetShortContent(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShortContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("GetShortContent() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
