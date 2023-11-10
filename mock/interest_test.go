package mock

import "testing"

func TestInterestPerYear(t *testing.T) {
	type args struct {
		year     int
		interest float64
		amount   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive Case",
			args: args{
				year: 5,
				interest: 2,
				amount: 1000000,
			},
			want: 1104080.8032,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterestPerYear(tt.args.year, tt.args.interest, tt.args.amount); got != tt.want {
				t.Errorf("InterestPerYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
