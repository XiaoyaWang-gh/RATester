package publisher

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestPublish_2(t *testing.T) {
	tests := []struct {
		name    string
		pub     DestinationPublisher
		desc    Descriptor
		wantErr bool
	}{
		{
			name:    "success",
			pub:     DestinationPublisher{fs: afero.NewMemMapFs()},
			desc:    Descriptor{TargetPath: "test.txt"},
			wantErr: false,
		},
		{
			name:    "error: no target path",
			pub:     DestinationPublisher{fs: afero.NewMemMapFs()},
			desc:    Descriptor{},
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

			p := tt.pub
			err := p.Publish(tt.desc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
