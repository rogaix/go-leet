package palindrom_number

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x == -1 {
		return false
	}

	last := x % 10
	if last == x {
		return true
	}

	var rev int
	rest := x
	for {
		rev = rev*10 + last
		if rev == x {
			return true
		}

		rest = rest / 10
		last = rest % 10
		if rest <= 0 {
			return false
		}
	}
}
