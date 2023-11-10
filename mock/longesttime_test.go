package mock

import "testing"

func TestLongestTime(t *testing.T) {
	type args struct {
		keyTime [][2]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive Case",
			args: args{
				keyTime: [][2]int{{0,2}, {1,9}, {0,12}, {2,15}},
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestTime(tt.args.keyTime); got != tt.want {
				t.Errorf("LongestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
