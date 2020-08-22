package protein

import (
	"errors"
)

var (
	// ErrInvalidBase blah
	ErrInvalidBase error = errors.New("InvalidBase error")
	//ErrStop blah
	ErrStop error = errors.New("Stop error")
)

// FromCodon converts from codon to protein.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA converts RNA to proteins.
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return proteins, nil
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}

	return proteins, nil
}
