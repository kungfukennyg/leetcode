package anagram

import (
	"strconv"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
