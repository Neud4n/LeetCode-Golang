package twosum

/*

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Constraints:
* 2 <= nums.lenght <= 10⁴
* -10⁹ <= nums[i] <= 10⁹
* -10⁹ <= target <= 10⁹
* Only one valid answer exists.

*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*func main() {

	// INPUT
	numbers1 := []int{2, 7, 11, 15}
	numbers2 := []int{3, 2, 4}
	numbers3 := []int{3, 3}
	numbers4 := []int{2, 5, 5, 11}

	fmt.Println(twoSum(numbers1, 9))
	fmt.Println("=============")
	fmt.Println(twoSum(numbers2, 6))
	fmt.Println("=============")
	fmt.Println(twoSum(numbers3, 6))
	fmt.Println("=============")
	fmt.Println(twoSum(numbers4, 10))

}*/
