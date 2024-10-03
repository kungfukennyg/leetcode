package parens

import (
	"strconv"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{"["},
			want: false,
		},
		{
			args: args{"()"},
			want: true,
		},
		{
			args: args{"()[]{}"},
			want: true,
		},
		{
			args: args{"(]"},
			want: false,
		},
		{
			args: args{"([])"},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
