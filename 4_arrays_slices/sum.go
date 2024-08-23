package array_slices

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lenNumbers := len(numbersToSum)
	// sum := make([]int, lenNumbers)

	// for i, numbers := range numbersToSum {
	// 	sum[i] = append(numbers)
	// }

	// return sum

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	// ! This will make it out of bound if slice is zero
	// var sums []int
	// for _, numbers := range numbersToSum {
	// 	// slice[low:high]
	// 	tail := numbers[1:]
	// 	sums = append(sums, Sum(tail))
	// }

	// return sums

	// ? We need to write a better logic for "safe" slice
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
