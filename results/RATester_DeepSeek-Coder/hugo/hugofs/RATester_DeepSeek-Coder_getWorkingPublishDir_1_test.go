package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestGetWorkingPublishDir_1(t *testing.T) {
	type args struct {
		cfg config.Provider
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
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

			got, got1 := getWorkingPublishDir(tt.args.cfg)
			if got != tt.want {
				t.Errorf("getWorkingPublishDir() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getWorkingPublishDir() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
