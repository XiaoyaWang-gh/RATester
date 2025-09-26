package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConfig_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantC   *LogConfig
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

			gotC, err := tt.p.GetConfig(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogProject.GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("LogProject.GetConfig() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
