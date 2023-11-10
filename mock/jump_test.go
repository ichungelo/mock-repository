package mock

import (
	"reflect"
	"testing"
)

func TestJump(t *testing.T) {
	type args struct {
		first  int
		second int
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive Case",
			args: args{
				first:  3,
				second: 6,
				length: 5,
			},
			want: []int{3, 6, 9, 12, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.args.first, tt.args.second, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
