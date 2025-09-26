package create

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestusesSiteVar_1(t *testing.T) {
	type args struct {
		fi hugofs.FileMetaInfo
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
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

			b := &contentBuilder{}
			got, err := b.usesSiteVar(tt.args.fi)
			if (err != nil) != tt.wantErr {
				t.Errorf("contentBuilder.usesSiteVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("contentBuilder.usesSiteVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
