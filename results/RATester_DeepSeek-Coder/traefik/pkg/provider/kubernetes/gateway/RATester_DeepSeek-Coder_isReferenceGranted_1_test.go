package gateway

import (
	"fmt"
	"testing"
)

func TestIsReferenceGranted_1(t *testing.T) {
	type args struct {
		fromKind      string
		fromNamespace string
		toGroup       string
		toKind        string
		toName        string
		toNamespace   string
	}
	tests := []struct {
		name    string
		p       *Provider
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

			if err := tt.p.isReferenceGranted(tt.args.fromKind, tt.args.fromNamespace, tt.args.toGroup, tt.args.toKind, tt.args.toName, tt.args.toNamespace); (err != nil) != tt.wantErr {
				t.Errorf("Provider.isReferenceGranted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
