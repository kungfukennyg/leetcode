package word

import (
	"fmt"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
