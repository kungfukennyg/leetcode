# Intuition
My initial approach to this solution was a naive loop that would subslice the input string upon finding the first character within the needle, returning if we successfully match every character within the subslice. This was simple enough to implement, but then I started reading about rolling hash functions and wondered if that would be any faster.

# Approach
My next two approaches utilize a polynomial rolling hash function to step through the input string once, computing the hash for the given window (len(needle)) until we hit a match or the end of the string. This approach works for most inputs, but does run into collisions with large inputs of similar characters. Collisions can be somewhat mitigated by increasing the `mod` value i.e. `1e11 + 7`, and this works for most test cases, but a few still fail. I tried two approaches from here; doing a final character-by-character comparison after finding a matching hash, and computing two different hashes and checking if they both match. Both approaches work to pass all test cases, but are they any faster?

# Benchmarks
There are benchmarks included below. The results show that both rolling hash approaches are slower in all cases than a more naive direct character comparison, on both large and small inputs. This shows that this approach, while interesting and fun to write, is more of a novelty than a practical approach, at least given the benchmark parameters I defined. 

```
goos: linux
goarch: amd64
pkg: example.com/leetcode/28-needle-haystack
cpu: 13th Gen Intel(R) Core(TM) i9-13900K
BenchmarkStrStr/singleHash-worst-case-0-10/5-32         	26061597	        45.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/doubleHash-worst-case-0-10/5-32         	24311079	        49.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/naive-worst-case-0-10/5-32              	235187823	         5.057 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/singleHash-worst-case-1-100/50-32       	 1512346	       774.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/doubleHash-worst-case-1-100/50-32       	 1464782	       812.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/naive-worst-case-1-100/50-32            	43860859	        27.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/singleHash-worst-case-2-1000/500-32     	  143881	      8369 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/doubleHash-worst-case-2-1000/500-32     	  137371	      8957 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/naive-worst-case-2-1000/500-32          	 5970483	       200.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/singleHash-random-match-0-10/5-32       	69236829	        17.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/doubleHash-random-match-0-10/5-32       	58448871	        20.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/naive-random-match-0-10/5-32            	218845314	         5.582 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/singleHash-random-match-1-100/50-32     	 2977251	       402.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/doubleHash-random-match-1-100/50-32     	 4256020	       283.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/naive-random-match-1-100/50-32          	40383946	        30.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/singleHash-random-match-2-1000/500-32   	  282153	      4143 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/doubleHash-random-match-2-1000/500-32   	  431841	      2779 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrStr/naive-random-match-2-1000/500-32        	 4365021	       279.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	example.com/leetcode/28-needle-haystack	25.971s


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


```

# Complexity
- Time complexity:
O(n + m) where n is the haystack length and m is the needle length

- Space complexity:
O(1) not counting the needle and haystack themselves

# Code
```
type hasher struct {
	base int
	mod  int
}

func (h *hasher) hashStr(s string, hash int) int {
	for i := 0; i < len(s); i++ {
		hash = (hash*h.base + int(s[i])) % h.mod
	}
	return hash
}

func (h *hasher) hash(b byte, hash int) int {
	return (hash*h.base + int(b)) % h.mod
}

func (h *hasher) negate(b byte, power, curHash int) int {
	curHash = (curHash - power*int(b)) % h.mod
	if curHash < 0 {
		curHash += h.mod
	}
	return curHash
}

func strStr(haystack string, needle string) int {
	window := len(needle)
	if window < 1 || window > len(haystack) {
		return -1
	}

	if window == 1 && len(haystack) == 1 && needle[0] == haystack[0] {
		return 0
	}

	hasher := hasher{base: 256, mod: 1e7 + 7}

	// compute the hash of our needle
	needleHash := hasher.hashStr(needle, 0)

	// compute the first hash window
	curHash := hasher.hashStr(haystack[0:window], 0)

	// short-circuit, first window matches
	// need to ensure the substrings actually match due to collisions
	if curHash == needleHash && stringsMatch(haystack[0:window], needle) {
		return 0
	}

	// compute base^(window-1)
	power := 1
	for i := 0; i < window-1; i++ {
		power = (power * hasher.base) % hasher.mod
	}

	// compute sliding window of hashes
	for i := 1; i <= len(haystack)-window; i++ {
		// remove effect of prev first char in window
		curHash = hasher.negate(haystack[i-1], power, curHash)
		// compute hash in new window
		curHash = hasher.hash(haystack[i+window-1], curHash)
		if curHash == needleHash && stringsMatch(haystack[i:i+window], needle) {
			return i
		}
	}
	return -1
}

func stringsMatch(s1 string, s2 string) bool {
	for i, s := range s1 {
		if s2[i] != byte(s) {
			return false
		}
	}
	return true
}

func strStrDoubleHash(haystack string, needle string) int {
	window := len(needle)
	if window < 1 || window > len(haystack) {
		return -1
	}

	if window == 1 && len(haystack) == 1 && needle[0] == haystack[0] {
		return 0
	}

	// two hashers to reduce likelihood of collisions
	// if still getting collisions, mod exponent can be increased
	hasherA := hasher{base: 256, mod: 1e7 + 7}
	hasherB := hasher{base: 31, mod: 1e9 + 7}

	// compute the hash of our needle
	needleHashA := hasherA.hashStr(needle, 0)
	needleHashB := hasherB.hashStr(needle, 0)

	// compute the first hash window
	curHashA := hasherA.hashStr(haystack[0:window], 0)
	curHashB := hasherB.hashStr(haystack[0:window], 0)

	// short-circuit, first window matches
	if curHashA == needleHashA && curHashB == needleHashB {
		return 0
	}

	// compute base^(window-1)
	powerA, powerB := 1, 1
	for i := 0; i < window-1; i++ {
		powerA = (powerA * hasherA.base) % hasherA.mod
		powerB = (powerB * hasherB.base) % hasherB.mod
	}

	// compute sliding window of hashes
	for i := 1; i <= len(haystack)-window; i++ {
		// remove effect of prev first char in window
		curHashA = hasherA.negate(haystack[i-1], powerA, curHashA)
		curHashB = hasherB.negate(haystack[i-1], powerB, curHashB)
		// compute hash in new window
		curHashA = hasherA.hash(haystack[i+window-1], curHashA)
		curHashB = hasherB.hash(haystack[i+window-1], curHashB)

		if curHashA == needleHashA && curHashB == needleHashB {
			return i
		}
	}
	return -1
}

func strStrNaive(haystack string, needle string) int {
	if len(needle) < 1 ||
		len(needle) > len(haystack) {
		return -1
	}

OUTER:
	for i, s := range haystack {
		if i+len(needle) > len(haystack) {
			break
		}
		if s != rune(needle[0]) {
			continue
		}

		sub := haystack[i : i+len(needle)]
		for j, ss := range sub {
			if needle[j] != byte(ss) {
				continue OUTER
			}
		}
		return i
	}

	return -1
}

```