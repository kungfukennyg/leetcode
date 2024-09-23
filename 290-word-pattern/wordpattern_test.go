package wordpattern

import (
	"strconv"
	"testing"
)

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				pattern: "abba",
				s:       "dog dog dog dog",
			},
			want: false,
		},
		{
			args: args{
				pattern: "aba",
				s:       "dog cat cat",
			},
			want: false,
		},
		{
			args: args{
				pattern: "aba",
				s:       "cat cat cat dog",
			},
			want: false,
		},
		{
			args: args{
				pattern: "abba",
				s:       "dog cat cat dog",
			},
			want: true,
		},
		{
			args: args{
				pattern: "abba",
				s:       "dog cat cat fish",
			},
			want: false,
		},
		{
			args: args{
				pattern: "aaaa",
				s:       "dog cat cat dog",
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
