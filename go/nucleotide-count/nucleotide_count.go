package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (dna DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, r := range dna {
		if _, found := h[r]; !found {
			return nil, fmt.Errorf("invalid nucleotide: %v", r)
		}
		h[r]++
	}

	return h, nil
}
