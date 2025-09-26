package hugio

import (
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	type fields struct {
		Patterns []*HasBytesPattern
		i        int
		done     bool
		buff     []byte
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
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

			h := &HasBytesWriter{
				Patterns: tt.fields.Patterns,
				i:        tt.fields.i,
				done:     tt.fields.done,
				buff:     tt.fields.buff,
			}
			gotN, err := h.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HasBytesWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("HasBytesWriter.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
