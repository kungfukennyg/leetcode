package justify

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				words:    []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
				maxWidth: 16},
			want: []string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country"},
		},
		{
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20},
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Error("fullJustify() =")
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("%d =\n%+v\n%+v", i, tt.want[i], got[i])
					}
				}
			}
		})
	}
}
func Test_fullJustifyTwoPass(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				words:    []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
				maxWidth: 16},
			want: []string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country"},
		},
		{
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20},
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := fullJustifyTwoPass(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Error("fullJustify() =")
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("%d =\n%+v\n%+v", i, tt.want[i], got[i])
					}
				}
			}
		})
	}
}
func Test_fullJustifyThreePass(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				words:    []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
				maxWidth: 16},
			want: []string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country"},
		},
		{
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20},
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := fullJustifyThreePass(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Error("Test_fullJustifyThreePass() =")
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("%d =\n%+v\n%+v", i, tt.want[i], got[i])
					}
				}
			}
		})
	}
}

func BenchmarkFullJustify(b *testing.B) {
	input := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	widths := []int{15, 30, 50, 100}

	funcs := []struct {
		name string
		fn   func([]string, int) []string
	}{
		{
			name: "",
			fn:   fullJustify,
		},
		{
			name: "TwoPass",
			fn:   fullJustifyTwoPass,
		},
		{
			name: "ThreePass",
			fn:   fullJustifyThreePass,
		},
	}

	for _, tc := range funcs {
		for _, width := range widths {
			b.Run(fmt.Sprintf("%s-%d", tc.name, width), func(b *testing.B) {
				for range b.N {
					tc.fn(input, width)
				}
			})
		}
	}
}
