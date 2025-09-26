package task

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTask_1(t *testing.T) {
	type args struct {
		tname string
		spec  string
		f     TaskFunc
		opts  []Option
	}
	tests := []struct {
		name string
		args args
		want *Task
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

			if got := NewTask(tt.args.tname, tt.args.spec, tt.args.f, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
