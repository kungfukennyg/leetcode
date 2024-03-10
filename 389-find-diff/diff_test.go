package diff

import (
	"fmt"
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{s: "a", t: "aa"},
			want: 'a',
		},
		{
			args: args{s: "abcd", t: "abcde"},
			want: 'e',
		},
		{
			args: args{s: "", t: "y"},
			want: 'y',
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
