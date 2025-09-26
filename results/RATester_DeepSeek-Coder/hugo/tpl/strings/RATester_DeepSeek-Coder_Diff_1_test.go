package strings

import (
	"fmt"
	"testing"
)

func TestDiff_1(t *testing.T) {
	type args struct {
		oldname string
		old     any
		newname string
		new     any
	}
	tests := []struct {
		name    string
		ns      *Namespace
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestDiff",
			ns:   &Namespace{},
			args: args{
				oldname: "old",
				old:     "old value",
				newname: "new",
				new:     "new value",
			},
			want:    "+new\n-old\n+value\n-value\n",
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

			got, err := tt.ns.Diff(tt.args.oldname, tt.args.old, tt.args.newname, tt.args.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Diff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
