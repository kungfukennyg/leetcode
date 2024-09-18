package isomorph

import (
	"strconv"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
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
				s: "badc",
				t: "baba",
			},
			want: false,
		},
		{
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
