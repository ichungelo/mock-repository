package mock

import "testing"

func TestIntReverser(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Positive Case",
			args: args{
				num: 1234,
			},
			want: 4321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntReverser(tt.args.num); got != tt.want {
				t.Errorf("IntReverser() = %v, want %v", got, tt.want)
			}
		})
	}
}
