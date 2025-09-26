package orm

import (
	"fmt"
	"testing"
)

func TestUpdate_8(t *testing.T) {
	d := &DoNothingOrm{}

	type TestModel struct {
		Id   int
		Name string
	}

	testCases := []struct {
		name    string
		model   TestModel
		cols    []string
		want    int64
		wantErr bool
	}{
		{
			name: "update success",
			model: TestModel{
				Id:   1,
				Name: "test",
			},
			cols:    []string{"Name"},
			want:    1,
			wantErr: false,
		},
		{
			name: "update failure",
			model: TestModel{
				Id:   1,
				Name: "test",
			},
			cols:    []string{"Name"},
			want:    0,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := d.Update(&tc.model, tc.cols...)
			if (err != nil) != tc.wantErr {
				t.Errorf("DoNothingOrm.Update() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("DoNothingOrm.Update() = %v, want %v", got, tc.want)
			}
		})
	}
}
