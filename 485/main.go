package main

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	result := 0

	for _, num := range nums {
		if num == 1 {
			count++
			if count > result {
				result = count
			}
		} else {
			count = 0
		}
	}
	return result
}
