package zigzag

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			args: args{s: "PAYPALISHIRING", numRows: 4},
			want: "PINALSIGYAHRPI",
		},
		{
			args: args{s: "A", numRows: 1},
			want: "A",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
