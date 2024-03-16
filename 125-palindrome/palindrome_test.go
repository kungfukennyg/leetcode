package palindrome

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{"0P"},
			want: false,
		},
		{
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			args: args{"race a car"},
			want: false,
		},
		{
			args: args{" "},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
