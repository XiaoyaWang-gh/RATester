package releaser

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Testlogf_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "Test case 1",
			format: "Hello, %s",
			args:   []interface{}{"World"},
			want:   "Hello, World",
		},
		{
			name:   "Test case 2",
			format: "Number: %d",
			args:   []interface{}{42},
			want:   "Number: 42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			logf(tt.format, tt.args...)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			if got := string(out); got != tt.want {
				t.Errorf("logf() = %v, want %v", got, tt.want)
			}
		})
	}
}
