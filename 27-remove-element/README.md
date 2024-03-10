# 27. Remove Element
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array<p>
int val = ...; // Value to remove<p>
int[] expectedNums = [...]; // The expected answer with correct length.<p>
                            // It is sorted with no values equaling val.<p>

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;<p>
sort(nums, 0, k); // Sort the first k elements of nums<p>
for (int i = 0; i < actualLength; i++) {<p>
    assert nums[i] == expectedNums[i];<p>
}<p>
If all assertions pass, then your solution will be accepted.