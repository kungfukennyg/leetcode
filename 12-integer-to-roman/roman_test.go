package roman

import (
	"fmt"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{num: 20},
			want: "XX",
		},
		{
			args: args{num: 3},
			want: "III",
		},
		{
			args: args{num: 58},
			want: "LVIII",
		},
		{
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
