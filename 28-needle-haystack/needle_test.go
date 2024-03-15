package needle

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				haystack: "mississippi",
				needle:   "sippia",
			},
			want: -1,
		},
		{
			args: args{
				haystack: "mississippi",
				needle:   "sippi",
			},
			want: 6,
		},
		{
			args: args{
				haystack: "mississippi",
				needle:   "sipp",
			},
			want: 6,
		},
		{
			args: args{
				haystack: "ababcaababcaabc",
				needle:   "ababcaabc",
			},
			want: 6,
		},
		{
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			args: args{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
		{
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
		{
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStrNaive(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
		{
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
		{
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := strStrNaive(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	inputs := [][2]int{{10, 5}, {100, 50}, {1000, 500}}
	funcs := []struct {
		name string
		fn   func(string, string) int
	}{
		{
			name: "singleHash",
			fn:   strStr,
		},
		{
			name: "doubleHash",
			fn:   strStrDoubleHash,
		},
		{
			name: "naive",
			fn:   strStrNaive,
		},
	}

	for i, in := range inputs {
		// worst case, no match
		haystack := randomStr(in[0])
		needleBytes := []rune(haystack[:in[1]])
		needleBytes[0] = rune(haystack[0] - 'z')
		needle := string(needleBytes)
		for _, fn := range funcs {
			b.Run(fmt.Sprintf("%s-worst-case-%d-%d/%d", fn.name, i, in[0], in[1]), func(b *testing.B) {
				for range b.N {
					fn.fn(haystack, needle)
				}
			})
		}
	}

	for i, in := range inputs {
		haystack := randomStr(in[0])
		needle := randomSubset(haystack, in[1])
		for _, fn := range funcs {
			b.Run(fmt.Sprintf("%s-random-match-%d-%d/%d", fn.name, i, in[0], in[1]), func(b *testing.B) {
				for range b.N {
					fn.fn(haystack, needle)
				}
			})
		}
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randomStr(n int) string {
	out := make([]rune, n)
	for i := range n {
		out[i] = letters[rand.Intn(len(letters))]
	}
	return string(out)
}

func randomSubset(in string, window int) string {
	start := rand.Intn(len(in) - window)
	return in[start:window]
}
