package prefix

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{strs: []string{"ab", "a"}},
			want: "a",
		},
		{
			args: args{strs: []string{""}},
			want: "",
		},
		{
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
