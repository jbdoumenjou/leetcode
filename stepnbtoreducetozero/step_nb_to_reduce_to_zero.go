package stepnbtoreducetozero

// This is a simulation problem;
// all we have to do to arrive at the correct result is follow the instructions.
// Through this problem, you will learn how to identify if a number is even or odd
// (hint, think about the modulo operator from the previous problem).
// For those who are looking for an added challenge,
// you can try solving this problem using bitwise operations.
// Both methods are introduced in the video solution.
// We encourage you to try solving 1342.
// Number of Steps to Reduce a Number to Zero on your own first,
// then return here to watch the video solution.
//
// time complexity: O(log(n))
// space complexity: O(1)
func numberOfSteps(num int) int {
	var stepNb int
	for stepNb = 0; num > 0; stepNb++ {
		if num%2 == 0 {
			num /= 2
			continue
		}

		num--
	}

	return stepNb
}

func numberOfSteps_bit(num int) int {
	var stepNb int
	if num == 0 {
		return 0
	}

	for num > 0 {
		// is odd
		if num&1 == 1 {
			num -= 1
			stepNb++
		}

		num >>= 1
		stepNb++
	}

	return stepNb - 1
}
