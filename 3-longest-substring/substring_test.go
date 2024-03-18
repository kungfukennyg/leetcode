package substring

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"au"},
			want: 2,
		},
		{
			args: args{" "},
			want: 1,
		},
		{
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			args: args{""},
			want: 0,
		},
		{
			args: args{"bbbbb"},
			want: 1,
		},
		{
			args: args{"pwwkew"},
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
