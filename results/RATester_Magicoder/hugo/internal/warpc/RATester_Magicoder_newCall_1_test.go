package warpc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewCall_1(t *testing.T) {
	type args struct {
		q Message[int]
	}
	tests := []struct {
		name    string
		args    args
		want    *call[int, int]
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

			d := &dispatcher[int, int]{}
			got, err := d.newCall(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
