package logs

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetLogger_1(t *testing.T) {
	tests := []struct {
		name     string
		prefixes []string
		want     *log.Logger
	}{
		{
			name:     "Test case 1",
			prefixes: []string{"test"},
			want:     GetLogger("test"),
		},
		{
			name:     "Test case 2",
			prefixes: []string{"test2"},
			want:     GetLogger("test2"),
		},
		{
			name:     "Test case 3",
			prefixes: []string{"test3"},
			want:     GetLogger("test3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetLogger(tt.prefixes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
