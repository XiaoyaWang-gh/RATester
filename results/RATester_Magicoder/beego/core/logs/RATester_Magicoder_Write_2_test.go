package logs

import (
	"fmt"
	"testing"
)

func TestWrite_2(t *testing.T) {
	bl := &BeeLogger{
		msgChan: make(chan *LogMsg, 100),
	}

	tests := []struct {
		name    string
		bl      *BeeLogger
		args    []byte
		wantN   int
		wantErr bool
	}{
		{
			name:    "TestWrite_Empty",
			bl:      bl,
			args:    []byte{},
			wantN:   0,
			wantErr: false,
		},
		{
			name:    "TestWrite_NonEmpty",
			bl:      bl,
			args:    []byte("test"),
			wantN:   4,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotN, err := tt.bl.Write(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeeLogger.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("BeeLogger.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
