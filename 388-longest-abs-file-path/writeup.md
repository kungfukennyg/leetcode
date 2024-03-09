# Intuition
I visualized this problem as a tree structure with many branches, and decomposed it into 3 steps; parsing the input into a tree, performing a depth-first search for all file nodes, and then walking backwards from each file node to the root to find the longest path.

# Approach
See above

# Complexity
- Time complexity:
Parsing the tree is O(n) where n is the number of newline components to the input. 
Finding all files within the tree is O(n + M) where n is the number of nodes in the tree and M is the depth of the tree.
Finding the longest file is O(n) where n is the number of files found.

- Space complexity:
O(n + M) where n is the input length and M is the depth of the tree.

# Code
```
import (
	"fmt"
	"strings"
)

const pathSeparator = "/"

type node struct {
	children []*node
	parent   *node
	value    string
}

func (fs *node) isFile() bool {
	return strings.Contains(fs.value, ".")
}

// findFiles performs a depth-first search to find all nodes that contain a file
func (fs *node) findFiles() []*node {
	if fs == nil {
		return nil
	}
	s := stack[*node]{fs}
	files := []*node{}
	for len(s) > 0 {
		ss, nn := pop(s)
		s = ss
		node := *nn
		if node.isFile() {
			files = append(files, node)
		}

		for _, c := range node.children {
			s = push(s, c)
		}
	}

	return files
}

type stack[T any] []T

func pop[T any](s []T) (stack[T], *T) {
	if len(s) == 0 {
		return s, nil
	}
	t := s[len(s)-1]
	s = s[:len(s)-1]
	return s, &t
}

func push[T any](s []T, t T) stack[T] {
	s = append(s, t)
	return s
}

func lengthLongestPath(input string) int {
	fs := parse(input)
	files := fs.findFiles()

	if !strings.Contains(input, ".") {
		return 0
	}

	if !strings.Contains(input, "\n") {
		return len(input)
	}

	longest := ""
	for _, node := range files {
		n := node
		path := ""
		// this loop will stop before including the root path /, which is
		// intended behavior
		for n.parent != nil {
			path = fmt.Sprint(n.value, pathSeparator, path)
			n = n.parent
		}
		if len(path) > len(longest) {
			longest = path
		}
	}

	longest = longest[:len(longest)-1]
	if strings.HasPrefix(longest, pathSeparator) {
		longest = longest[1:]
	}
	return len(longest)
}

func parse(input string) *node {
	root := &node{value: "/", children: []*node{}}
	if len(input) == 1 {
		root.value = input
		return root
	}

	cur := &node{
		children: []*node{},
		parent:   root,
		value:    "",
	}
	root.children = append(root.children, cur)
	depth := 1
	for i, s := range strings.Split(input, "\n") {
		if i == 0 {
			cur.value = s
			depth = 1
			continue
		}

		value := strings.ReplaceAll(s, "\t", "")
		curDepth := strings.Count(s, "\t")
		if curDepth == depth {
			// if at right depth, this is a child of the current node
			cur.children = append(cur.children, &node{
				children: []*node{},
				parent:   cur,
				value:    value,
			})
		} else if curDepth > depth {
			// further down, this is a child of the last-added node
			depth++
			cur = cur.children[len(cur.children)-1]
			cur.children = append(cur.children, &node{
				children: []*node{},
				parent:   cur,
				value:    value,
			})
		} else if curDepth < depth {
			// higher up, walk back up the hierarchy
			for depth > curDepth {
				cur = cur.parent
				depth--
			}
			cur.children = append(cur.children, &node{
				children: []*node{},
				parent:   cur,
				value:    value,
			})
		}
	}
	return root
}


```

```
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

```