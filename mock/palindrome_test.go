package mock

import "testing"

func TestPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive Case",
			args: args{
				input: "ABBA",
			},
			want: true,
		},
		{
			name: "Negative Case (Not Palindrome)",
			args: args{
				input: "ABSA",
			},
			want: false,
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.args.input); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
