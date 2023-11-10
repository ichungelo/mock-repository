package mock

import "testing"

func TestContainWord(t *testing.T) {
	type args struct {
		word     string
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive Case",
			args: args{
				word: "jumps",
				sentence: "quick brown fox jumps over the lazy dog",
			},
			want: true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainWord(tt.args.word, tt.args.sentence); got != tt.want {
				t.Errorf("ContainWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
