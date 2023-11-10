package mock

import "testing"

func TestSumOfSeries(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive Case",
			args: args{
				n: 10,
			},
			want: 2.3166247903554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfSeries(tt.args.n); got != tt.want {
				t.Errorf("SumOfSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}
