// Package leap is a solution to exercism.io exercise titled Leap.
package leap

func isDivisibleBy4(number int) bool {
	if number%4 == 0 {
		return true
	}
	return false
}

func isDivisibleBy100(number int) bool {
	if number%100 == 0 {
		return true
	}
	return false
}

func isDivisibleBy400(number int) bool {
	if number%400 == 0 {
		return true
	}
	return false
}

// IsLeapYear reports if a given year is a leap year.
func IsLeapYear(year int) bool {
	if isDivisibleBy400(year) || !isDivisibleBy100(year) && isDivisibleBy4(year) {
		return true
	}
	return false

	// // Or cleaner
	// if year%400 == 0 || year%100 != 0 && year%4 == 0 {
	// 	return true
	// }
	// return false
}
