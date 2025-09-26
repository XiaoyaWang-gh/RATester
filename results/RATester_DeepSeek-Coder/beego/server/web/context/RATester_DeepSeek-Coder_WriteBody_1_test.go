package context

import (
	"fmt"
	"io"
	"testing"
)

func TestWriteBody_1(t *testing.T) {
	type args struct {
		encoding string
		writer   io.Writer
		content  []byte
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

			got, got1, err := WriteBody(tt.args.encoding, tt.args.writer, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteBody() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("WriteBody() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
