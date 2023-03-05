package sequences

import (
	"math"
	"sort"
)

// longestConsecutive
// Complexity Big O
// sort: O(n log n)
// loop: O(n)
// the complexity of the func is O(n log n) since is the biggest
// Example 1:
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	longestSeq := 1
	currSeq := 1
	for i, value := range nums[:len(nums)-1] {
		if value != nums[i+1] {
			if value == nums[i+1]-1 { //check if the current number extends the sequence
				currSeq++
			} else { // the sequence is broken, so we record our current sequence and reset it to 1
				longestSeq = int(math.Max(float64(longestSeq), float64(currSeq))) // choose the biggest
				currSeq = 1                                                       // reset counter
			}
		}
	}
	return int(math.Max(float64(longestSeq), float64(currSeq))) // because the last element can be part of the biggest sequence
}

// longestConsecutiveHash
func longestConsecutiveHash(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hash := make(map[int]bool)
	longestSeq := 1
	currSeq := 1
	for _, v := range nums {
		hash[v] = true
	}
	for v := range hash {
		if !hash[v-1] { // it means is the beginning of a new seq
			currSeq = 1
			for hash[v+1] { // while is a consecutive number of v the currseq increases
				currSeq++
				v++
			}
			if longestSeq < currSeq {
				longestSeq = currSeq
			}
		}
	}
	return int(math.Max(float64(longestSeq), float64(currSeq)))
}
