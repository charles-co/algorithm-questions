package randoms

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	for i := range ans {
		ans[i] = 1
	}

	suffix := 1
	prefix := 1

	for i := 0; i < len(nums); i++ {
		ans[i] *= prefix
		prefix *= nums[i]
		ans[len(ans)-i-1] *= suffix
		suffix *= nums[len(ans)-i-1]

	}
	return ans
}
