package mock

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
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
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prime(tt.args.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prime() = %v, want %v", got, tt.want)
			}
		})
	}
}
