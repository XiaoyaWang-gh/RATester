package file

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewJsonDb_1(t *testing.T) {
	type args struct {
		runPath string
	}
	tests := []struct {
		name string
		args args
		want *JsonDb
	}{
		{
			name: "TestNewJsonDb",
			args: args{
				runPath: "test",
			},
			want: &JsonDb{
				RunPath:        "test",
				TaskFilePath:   filepath.Join("test", "conf", "tasks.json"),
				HostFilePath:   filepath.Join("test", "conf", "hosts.json"),
				ClientFilePath: filepath.Join("test", "conf", "clients.json"),
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

			got := NewJsonDb(tt.args.runPath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJsonDb() = %v, want %v", got, tt.want)
			}
		})
	}
}
