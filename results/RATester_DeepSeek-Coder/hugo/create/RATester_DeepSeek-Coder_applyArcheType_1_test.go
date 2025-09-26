package create

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestApplyArcheType_1(t *testing.T) {
	type args struct {
		contentFilename string
		archetypeFi     hugofs.FileMetaInfo
	}
	tests := []struct {
		name    string
		b       *contentBuilder
		args    args
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

			if err := tt.b.applyArcheType(tt.args.contentFilename, tt.args.archetypeFi); (err != nil) != tt.wantErr {
				t.Errorf("contentBuilder.applyArcheType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
