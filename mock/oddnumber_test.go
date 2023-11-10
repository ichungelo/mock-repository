package mock

import (
	"reflect"
	"testing"
)

func TestOddNumber(t *testing.T) {
	type args struct {
		total int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive Case",
			args: args{
				total: 10,
			},
			want: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddNumber(tt.args.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
