package mock

import "testing"

func TestCloseBracket(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Positive Case",
			args: args{
				input: "[{}]({[]})",
			},
			want: true,
		},
		{
			name: "Start With Close Bracket",
			args: args{
				input: "}}{{{}}}",
			},
			want: false,
		},
		{
			name: "Bracket Not Close",
			args: args{
				input: "{{[[]}}",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloseBracket(tt.args.input); got != tt.want {
				t.Errorf("CloseBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}
