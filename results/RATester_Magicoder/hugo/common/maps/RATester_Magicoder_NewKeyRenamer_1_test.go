package maps

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestNewKeyRenamer_1(t *testing.T) {
	type args struct {
		patternKeys []string
	}
	tests := []struct {
		name    string
		args    args
		want    KeyRenamer
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				patternKeys: []string{"key1", "newKey1", "key2", "newKey2"},
			},
			want: KeyRenamer{
				renames: []keyRename{
					{
						pattern: glob.MustCompile("key1", '/'),
						newKey:  "newKey1",
					},
					{
						pattern: glob.MustCompile("key2", '/'),
						newKey:  "newKey2",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				patternKeys: []string{"key1", "newKey1", "invalidPattern", "newKey2"},
			},
			want:    KeyRenamer{},
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

			got, err := NewKeyRenamer(tt.args.patternKeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKeyRenamer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyRenamer() = %v, want %v", got, tt.want)
			}
		})
	}
}
