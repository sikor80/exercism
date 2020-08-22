package strain

// Ints int slice.
type Ints []int

// Lists Ints slice.
type Lists []Ints

// Strings string slice.
type Strings []string

// Keep keep all elements in Ints which eval predicate function is true.
func (i Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, item := range i {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Discard discard all elements in Ints which eval pred function is false
func (i Ints) Discard(predicate func(int) bool) Ints {
	return i.Keep(func(x int) bool {
		return !predicate(x)
	})
}

// Keep keep all elements in Lists which eval pred function is true
func (l Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, item := range l {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Discard discard all elements in Lists which eval pred function is false
func (l Lists) Discard(predicate func([]int) bool) Lists {
	return l.Keep(func(x []int) bool {
		return !predicate(x)
	})
}

// Keep keep all elements in Strings which eval pred function is true
func (s Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, item := range s {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Discard discard all elements in Strings which eval pred function is false
func (s Strings) Discard(predicate func(string) bool) Strings {
	return s.Keep(func(x string) bool {
		return !predicate(x)
	})
}
