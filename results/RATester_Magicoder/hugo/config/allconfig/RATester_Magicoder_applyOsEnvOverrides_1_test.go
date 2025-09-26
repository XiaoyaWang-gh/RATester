package allconfig

import (
	"fmt"
	"testing"
)

func TestapplyOsEnvOverrides_1(t *testing.T) {
	type args struct {
		environ []string
	}
	tests := []struct {
		name    string
		l       configLoader
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			l:    configLoader{},
			args: args{
				environ: []string{"HUGO_TEST_FIELD=value"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			l:    configLoader{},
			args: args{
				environ: []string{"HUGO_TEST_FIELD=invalid_value"},
			},
			wantErr: true,
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

			if err := tt.l.applyOsEnvOverrides(tt.args.environ); (err != nil) != tt.wantErr {
				t.Errorf("applyOsEnvOverrides() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
