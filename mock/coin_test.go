package mock

import "testing"

func TestCoin(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive Case",
			args: args{
				n: 4,
				k: 5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Coin(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("Coin() = %v, want %v", got, tt.want)
			}
		})
	}
}
