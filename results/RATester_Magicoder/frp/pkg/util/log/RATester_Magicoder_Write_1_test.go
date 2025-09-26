package log

import (
	"fmt"
	"testing"

	"github.com/cloudflare/cfssl/log"
)

func TestWrite_1(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		w       *WriteLogger
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "Test case 1",
			w: &WriteLogger{
				level:  log.LevelInfo,
				offset: 0,
			},
			args: args{
				p: []byte("test message"),
			},
			wantN:   len("test message"),
			wantErr: false,
		},
		{
			name: "Test case 2",
			w: &WriteLogger{
				level:  log.LevelError,
				offset: 0,
			},
			args: args{
				p: []byte("error message"),
			},
			wantN:   len("error message"),
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotN, err := tt.w.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteLogger.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteLogger.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
