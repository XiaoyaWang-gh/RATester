package cert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate_1(t *testing.T) {
	type args struct {
		commonName string
	}
	tests := []struct {
		name    string
		args    args
		want    *Artifacts
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				commonName: "test.com",
			},
			want: &Artifacts{
				Key:             []byte("test key"),
				Cert:            []byte("test cert"),
				CAKey:           []byte("test CA key"),
				CACert:          []byte("test CA cert"),
				ResourceVersion: "1",
			},
			wantErr: false,
		},
		// add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cp := &SelfSignedCertGenerator{}
			got, err := cp.Generate(tt.args.commonName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
