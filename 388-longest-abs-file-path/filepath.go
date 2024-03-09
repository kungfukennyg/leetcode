package filepath

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
