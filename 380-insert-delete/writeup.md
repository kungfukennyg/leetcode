# Intuition
I approached this problem knowing I wanted to use a map-backed set for the constant-time insertions and deletions. My first idea was to use Golang's seemingly-random map iteration order to provide random values, but this is neither a novel idea [nor suitable for a uniformly-distributed range of values](https://dev.to/wallyqs/gos-map-iteration-order-is-not-that-random-mag). It also wouldn't be ideal to rely on an implementation detail of Golang that could change even if it did provide uniformly-distributed values now. After all, the entire point of Golang not providing consistent map iteration over keys is to prevent people from relying on that behavior.

# Approach
A set provides constant-time insertion and deletion, but does not provide an easy manner on its own to obtain uniformly-distributed random keys without needing to iterate through all values. To provide random access two further datastructures are utilized; a slice of all of the map keys, and a map of values to their respective positions in the slice of keys. The slice of keys provides constant random access time, and the map of indices provides constant time deletions from the keys.

# Complexity
- Time complexity:
Insert: O(1) - Map insertion is O(1) and slice append is amortized to O(1) but can be O(n) in the worst case.
Delete: O(1) - Map key removal is O(1) and trimming one slice element by shuffling it to the back of the slice is O(1)

- Space complexity:
O(n) -- There are three backing collections each storing n elements, making the overall complexity O(3n) which is amortized to O(n). This could be reduced further by collapsing the indices map into the set rather than storing separately, but is minimal regardless.