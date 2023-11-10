package mock

import "testing"

func TestMinimumGroup(t *testing.T) {
	type args struct {
		predators []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive Case",
			args: args{
				predators: []int{-1, 8, 6, 0, 7, 3, 8, 9, -1, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumGroup(tt.args.predators); got != tt.want {
				t.Errorf("MinimumGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
