package filepath

import (
	"fmt"
	"testing"
)

func Test_lengthLongestPath(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"file1.txt\nfile2.txt\nlongfile.txt"},
			want: 12,
		},
		{
			args: args{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"},
			want: 20,
		},
		{
			args: args{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"},
			want: 32,
		},
		{
			args: args{"a"},
			want: 0,
		},
		{
			args: args{"a.txt"},
			want: 5,
		},
		{
			args: args{"a\n\tb.txt\na2\n\tb2.txt"},
			want: 9,
		},
		{
			args: args{"a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png"},
			want: 29,
		},
		{
			args: args{"a\n\tb\n\t\tc\n\t\t\td\n\t\t\t\te.txt\n\t\t\t\talsdkjf.txt\n\t\tskdjfl.txtlsdkjflsdjflsajdflkjasklfjkasljfklas\n\tlskdjflkajsflj.txt"},
			want: 50,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Logf("\n%s", tt.args.input)
			if got := lengthLongestPath(tt.args.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
