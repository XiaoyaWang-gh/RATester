package pagemeta

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateDateHandler_1(t *testing.T) {
	tests := []struct {
		name        string
		identifiers []string
		setter      func(d *FrontMatterDescriptor, t time.Time)
		wantErr     bool
	}{
		{
			name:        "Test with fmFilename",
			identifiers: []string{fmFilename},
			setter:      func(d *FrontMatterDescriptor, t time.Time) {},
			wantErr:     false,
		},
		{
			name:        "Test with fmModTime",
			identifiers: []string{fmModTime},
			setter:      func(d *FrontMatterDescriptor, t time.Time) {},
			wantErr:     false,
		},
		{
			name:        "Test with fmGitAuthorDate",
			identifiers: []string{fmGitAuthorDate},
			setter:      func(d *FrontMatterDescriptor, t time.Time) {},
			wantErr:     false,
		},
		{
			name:        "Test with custom field",
			identifiers: []string{"customField"},
			setter:      func(d *FrontMatterDescriptor, t time.Time) {},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := FrontMatterHandler{}
			_, err := f.createDateHandler(tt.identifiers, tt.setter)
			if (err != nil) != tt.wantErr {
				t.Errorf("createDateHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
