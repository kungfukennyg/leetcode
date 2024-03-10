# Intuition
My initial approach was to create two maps and count the number of letters in each string, and find where their counts did not match. When this proved slower than I liked I pivoted to allocating a byte slice for each string and modifying them in place. To accomplish this I treat the `t` slice as a buffer, popping each value off the top and checking if it could be found in `s`. When found, `t` is sub-sliced by one and an in-place rewrite is performed on `s` to mark that letter as used.   
# Approach
See above.

# Complexity
- Time complexity:
O(n^2) in the worst case

- Space complexity:
O(n)

# Code
```
func findTheDifference(s, t string) byte {
	a, b := []byte(s), []byte(t)
OUTER:
	for len(b) > 0 {
		s := b[0]
		for i, v := range a {
			if s == byte(v) {
				last := a[len(a)-1]
				a[i] = last
				a = a[:len(a)-1]
				b = b[1:]
				continue OUTER
			}
		}
		return s
	}

	return 0
}
```