package groupanagrams

import (
	"strconv"
	"strings"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		args args
		want [][]string
	}{
		{
			args: args{[]string{"", ""}},
			want: [][]string{{""}, {""}},
		},
		{
			args: args{
				strs: []string{
					"eat", "tea", "tan", "ate", "nat", "bat",
				},
			},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			args: args{[]string{""}},
			want: [][]string{{""}},
		},
		{
			args: args{[]string{"a"}},
			want: [][]string{{"a"}},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			if len(got) != len(tt.want) {
				t.Fatalf("groupAnagrams() = %v, want %v", got, tt.want)
			}
			wordCount := map[string]int{}
			for _, g := range got {
				wordCount[strings.Join(g, " ")] += 1
			}
			for _, w := range tt.want {
				str := strings.Join(w, " ")
				wordCount[str]--
				if wordCount[str] < 0 {
					t.Fatalf("groupAnagrams() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
