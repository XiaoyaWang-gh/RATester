package context

import (
	"fmt"
	"io"
	"testing"
)

func TestwriteLevel_1(t *testing.T) {
	type args struct {
		encoding string
		writer   io.Writer
		reader   io.Reader
		level    int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   string
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

			got, got1, err := writeLevel(tt.args.encoding, tt.args.writer, tt.args.reader, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("writeLevel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("writeLevel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
