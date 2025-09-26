package cert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetCA_1(t *testing.T) {
	type args struct {
		caKey  []byte
		caCert []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSetCA",
			args: args{
				caKey:  []byte("testKey"),
				caCert: []byte("testCert"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cp := &SelfSignedCertGenerator{}
			cp.SetCA(tt.args.caKey, tt.args.caCert)
			if !reflect.DeepEqual(cp.caKey, tt.args.caKey) || !reflect.DeepEqual(cp.caCert, tt.args.caCert) {
				t.Errorf("SetCA() = %v, want %v", cp, tt.args)
			}
		})
	}
}
