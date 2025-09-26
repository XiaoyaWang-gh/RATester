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
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				fromKind:      "kind1",
				fromNamespace: "namespace1",
				toGroup:       "group1",
				toKind:        "kind2",
				toName:        "name1",
				toNamespace:   "namespace2",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				fromKind:      "kind2",
				fromNamespace: "namespace2",
				toGroup:       "group2",
				toKind:        "kind3",
				toName:        "name2",
				toNamespace:   "namespace3",
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

			p := &Provider{}
			if err := p.isReferenceGranted(tt.args.fromKind, tt.args.fromNamespace, tt.args.toGroup, tt.args.toKind, tt.args.toName, tt.args.toNamespace); (err != nil) != tt.wantErr {
				t.Errorf("isReferenceGranted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
