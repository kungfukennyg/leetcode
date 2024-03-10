# 384. Shuffle an Array
Given an integer array nums, design an algorithm to randomly shuffle the array. All permutations of the array should be equally likely as a result of the shuffling.

Implement the Solution class:

Solution(int[] nums) Initializes the object with the integer array nums.<p>
int[] reset() Resets the array to its original configuration and returns it.<p>
int[] shuffle() Returns a random shuffling of the array.
 

## Example 1:

Input<p>
["Solution", "shuffle", "reset", "shuffle"]<p>
[[[1, 2, 3]], [], [], []]<p>
Output<p>
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]

## Constraints:

- 1 <= nums.length <= 50
- -106 <= nums[i] <= 106
- All the elements of nums are unique.
- At most 104 calls in total will be made to reset and shuffle.