package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_4(t *testing.T) {
	type fields struct {
		writers       [LevelDebug + 1 + 1]*fileLogWriter
		fullLogWriter *fileLogWriter
		Separate      []string
	}
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name    string
		fields  fields
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

			f := &multiFileLogWriter{
				writers:       tt.fields.writers,
				fullLogWriter: tt.fields.fullLogWriter,
				Separate:      tt.fields.Separate,
			}
			if err := f.WriteMsg(tt.args.lm); (err != nil) != tt.wantErr {
				t.Errorf("multiFileLogWriter.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
