# 6. Zigzag Conversion
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N<p>
A P L S I I G<p>
Y   I   R<p>
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

```string convert(string s, int numRows);```
 

## Example 1:

Input: s = "PAYPALISHIRING", numRows = 3<p>
Output: "PAHNAPLSIIGYIR"
## Example 2:

Input: s = "PAYPALISHIRING", numRows = 4<p>
Output: "PINALSIGYAHRPI"
## Explanation:
<pre>
P     I    N
A   L S  I G
Y A   H R
P     I
</pre>
# Example 3:

Input: s = "A", numRows = 1<p>
Output: "A"
 

## Constraints:

- 1 <= s.length <= 1000
- s consists of English letters (lower-case and upper-case), ',' and '.'.
- 1 <= numRows <= 1000