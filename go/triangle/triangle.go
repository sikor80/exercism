// Package triangle is a solution to exercism.io exercise titled Triangle.
package triangle

import "math"

// Kind represents a kind of triangle.
type Kind int

const (
	// NaT not a triangle.
	NaT Kind = iota
	// Equ equilateral.
	Equ
	// Iso isosceles.
	Iso
	// Sca scalene.
	Sca
)

// KindFromSides generates kind of trianlge given the leghts of sides.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	hasZero := a <= 0 || b <= 0 || c <= 0
	failedSumCheck := a+b < c || a+c < b || b+c < a
	hasNaN := math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
	hasInf := math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)

	switch {
	case hasZero || failedSumCheck || hasNaN || hasInf:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	case a != b && b != c:
		k = Sca
	default:
		k = NaT
	}

	return k
}
