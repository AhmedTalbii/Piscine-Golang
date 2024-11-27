// Two Sum brut force

package main 

import "fmt"

// func twoSum(nums []int, target int) []int {
//     for i:=0 ; i < len(nums) ;i++ {
// 		for j:=i+1 ; j < len(nums) ;j++ {
// 			if nums[i] + nums[j] == target  {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
    // Create a map to store the values and their indices
    seen := make(map[int]int)

    // Loop through the array
    for i, num := range nums {
        // Calculate the complement
        complement := target - num

		
		fmt.Println(i," ",num," ",seen[complement])
        // Check if the complement is in the map
        if index, exists := seen[complement]; exists {
           
            // Return the indices if a match is found

            return []int{index, i}
        }

        // Store the current number and its index in the map
        seen[num] = i
    }

    // Return nil if no solution is found
    return nil
}


func main() {
    nums := []int{9,3,6,8}
	target := 9

	result := twoSum(nums, target)
    fmt.Println(result)
	
}