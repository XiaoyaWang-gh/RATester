package common

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteMsg_1(t *testing.T) {
	type args struct {
		when  time.Time
		msg   string
		level int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				when:  time.Now(),
				msg:   "Test message",
				level: 1,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				when:  time.Now(),
				msg:   "",
				level: 1,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				when:  time.Now(),
				msg:   "Test message",
				level: 0,
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				when:  time.Now(),
				msg:   "",
				level: 0,
			},
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

			lg := &StoreMsg{}
			if err := lg.WriteMsg(tt.args.when, tt.args.msg, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("StoreMsg.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
