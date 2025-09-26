package orm

import (
	"fmt"
	"sync"
	"testing"
)

func TestAdd_4(t *testing.T) {
	type args struct {
		name string
		al   *alias
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test add",
			args: args{
				name: "test",
				al: &alias{
					Name: "test",
				},
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

			ac := &_dbCache{
				mux:   sync.RWMutex{},
				cache: make(map[string]*alias),
			}
			if got := ac.add(tt.args.name, tt.args.al); got != tt.want {
				t.Errorf("_dbCache.add() = %v, want %v", got, tt.want)
			}
		})
	}
}
