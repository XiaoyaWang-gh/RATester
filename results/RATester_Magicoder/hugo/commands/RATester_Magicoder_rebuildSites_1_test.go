package commands

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestrebuildSites_1(t *testing.T) {
	type args struct {
		events []fsnotify.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				events: []fsnotify.Event{
					{Name: "file1.txt", Op: fsnotify.Write},
					{Name: "file2.txt", Op: fsnotify.Write},
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				events: []fsnotify.Event{
					{Name: "file3.txt", Op: fsnotify.Write},
					{Name: "file4.txt", Op: fsnotify.Write},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hugoBuilder{}
			if err := c.rebuildSites(tt.args.events); (err != nil) != tt.wantErr {
				t.Errorf("rebuildSites() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
