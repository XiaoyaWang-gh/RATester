package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadYmlReader_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantCnf map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				path: "test.yml",
			},
			wantCnf: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				path: "nonexistent.yml",
			},
			wantCnf: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotCnf, err := ReadYmlReader(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadYmlReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCnf, tt.wantCnf) {
				t.Errorf("ReadYmlReader() = %v, want %v", gotCnf, tt.wantCnf)
			}
		})
	}
}
