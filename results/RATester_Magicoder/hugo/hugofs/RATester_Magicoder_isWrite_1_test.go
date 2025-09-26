package hugofs

import (
	"fmt"
	"os"
	"testing"
)

func TestisWrite_1(t *testing.T) {
	tests := []struct {
		name string
		flag int
		want bool
	}{
		{"Test1", os.O_RDWR, true},
		{"Test2", os.O_WRONLY, true},
		{"Test3", os.O_RDONLY, false},
		{"Test4", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isWrite(tt.flag); got != tt.want {
				t.Errorf("isWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
