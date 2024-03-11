package roman

import (
	"fmt"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{s: "III"},
			want: 3,
		},
		{
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
