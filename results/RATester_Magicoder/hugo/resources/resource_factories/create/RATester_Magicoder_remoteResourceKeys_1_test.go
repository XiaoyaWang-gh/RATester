package create

import (
	"fmt"
	"testing"
)

func TestremoteResourceKeys_1(t *testing.T) {
	type args struct {
		uri     string
		options map[string]any
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

			got, got1 := remoteResourceKeys(tt.args.uri, tt.args.options)
			if got != tt.want {
				t.Errorf("remoteResourceKeys() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("remoteResourceKeys() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
