package cache

import (
	"fmt"
	"testing"
)

func TestStartAndGC_1(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		bc      *MemoryCache
		args    args
		wantErr bool
	}{
		{
			name: "TestStartAndGC_0",
			bc:   &MemoryCache{},
			args: args{
				config: `{"interval": 10}`,
			},
			wantErr: false,
		},
		{
			name: "TestStartAndGC_1",
			bc:   &MemoryCache{},
			args: args{
				config: `{"interval": "a"}`,
			},
			wantErr: true,
		},
		{
			name: "TestStartAndGC_2",
			bc:   &MemoryCache{},
			args: args{
				config: `{}`,
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

			bc := tt.bc
			err := bc.StartAndGC(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartAndGC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
