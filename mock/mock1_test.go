package mock

import (
	"reflect"
	"testing"
)

func TestGetNumber(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		// TODO: Add test cases.
		{
			name: "Positive Test",
			args: args{
				arr:    []int{1, 2, 3, 2, 3, 4, 5},
				target: 6,
			},
			want: [2]int{3, 3},
		},
		{
			name: "No Match",
			args: args{
				arr:    []int{1, 2, 3, 2, 3, 4, 5},
				target: 10,
			},
			want: [2]int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumber(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
