package alils

import (
	"fmt"
	"testing"
)

func TestDeleteLogStore_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		p       *LogProject
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

			p := &LogProject{
				Name:            tt.p.Name,
				Endpoint:        tt.p.Endpoint,
				AccessKeyID:     tt.p.AccessKeyID,
				AccessKeySecret: tt.p.AccessKeySecret,
			}
			if err := p.DeleteLogStore(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("LogProject.DeleteLogStore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
