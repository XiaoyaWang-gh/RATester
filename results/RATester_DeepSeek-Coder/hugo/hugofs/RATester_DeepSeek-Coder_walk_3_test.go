package hugofs

import (
	"fmt"
	"testing"
)

func TestWalk_3(t *testing.T) {
	type args struct {
		path       string
		info       FileMetaInfo
		dirEntries []FileMetaInfo
	}
	tests := []struct {
		name    string
		w       *Walkway
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

			w := &Walkway{
				logger: tt.w.logger,
				walked: tt.w.walked,
				cfg:    tt.w.cfg,
			}
			if err := w.walk(tt.args.path, tt.args.info, tt.args.dirEntries); (err != nil) != tt.wantErr {
				t.Errorf("Walkway.walk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
