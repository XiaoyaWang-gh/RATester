package nathole

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewController_1(t *testing.T) {
	type args struct {
		analysisDataReserveDuration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *Controller
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				analysisDataReserveDuration: 1 * time.Second,
			},
			want: &Controller{
				clientCfgs: make(map[string]*ClientCfg),
				sessions:   make(map[string]*Session),
				analyzer:   NewAnalyzer(1 * time.Second),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				analysisDataReserveDuration: 2 * time.Second,
			},
			want: &Controller{
				clientCfgs: make(map[string]*ClientCfg),
				sessions:   make(map[string]*Session),
				analyzer:   NewAnalyzer(2 * time.Second),
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

			got, err := NewController(tt.args.analysisDataReserveDuration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewController() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
