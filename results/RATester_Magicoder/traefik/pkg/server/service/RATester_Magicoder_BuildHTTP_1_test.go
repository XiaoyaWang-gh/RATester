package service

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBuildHTTP_1(t *testing.T) {
	type args struct {
		rootCtx     context.Context
		serviceName string
	}
	tests := []struct {
		name    string
		args    args
		want    http.Handler
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

			m := &Manager{}
			got, err := m.BuildHTTP(tt.args.rootCtx, tt.args.serviceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.BuildHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.BuildHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
