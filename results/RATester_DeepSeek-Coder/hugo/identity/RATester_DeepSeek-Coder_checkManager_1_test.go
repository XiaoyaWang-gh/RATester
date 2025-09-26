package identity

import (
	"fmt"
	"testing"
)

func TestCheckManager_1(t *testing.T) {
	type args struct {
		sid   *searchID
		m     Manager
		level int
	}
	tests := []struct {
		name string
		args args
		want FinderResult
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

			f := &Finder{}
			if got := f.checkManager(tt.args.sid, tt.args.m, tt.args.level); got != tt.want {
				t.Errorf("Finder.checkManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
