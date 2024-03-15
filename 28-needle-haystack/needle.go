package needle

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

	for i, s := range haystack {
		if i+len(needle) > len(haystack) {
			break
		}
		if s != rune(needle[0]) {
			continue
		}

		sub := haystack[i : i+len(needle)]
		if stringsMatch(sub, needle) {
			return i
		}
	}

	return -1
}
