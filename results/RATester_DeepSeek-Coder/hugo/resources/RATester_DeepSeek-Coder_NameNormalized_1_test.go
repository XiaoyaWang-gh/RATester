package resources

import (
	"fmt"
	"testing"
)

func TestNameNormalized_1(t *testing.T) {
	type fields struct {
		sd ResourceSourceDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				sd: ResourceSourceDescriptor{
					NameNormalized: "test_name_normalized",
				},
			},
			want: "test_name_normalized",
		},
		{
			name: "Test case 2",
			fields: fields{
				sd: ResourceSourceDescriptor{
					NameNormalized: "test_name_normalized_2",
				},
			},
			want: "test_name_normalized_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &genericResource{
				sd: tt.fields.sd,
			}
			if got := l.NameNormalized(); got != tt.want {
				t.Errorf("genericResource.NameNormalized() = %v, want %v", got, tt.want)
			}
		})
	}
}
