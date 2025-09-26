package filenotify

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew_8(t *testing.T) {
	type args struct {
		interval time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    FileWatcher
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				interval: 1 * time.Second,
			},
			want:    nil, // This should be the actual FileWatcher instance, but it's not possible to create one here
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := New(tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
