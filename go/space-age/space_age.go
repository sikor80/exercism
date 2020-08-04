// Package space is a solution to exercism.io exercise titled Space Age.
package space

const (
	// Earth orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
	Earth = 31557600
)

// Planet represents planet object.
type Planet string

// Age calculates how old someone would be on a given plant.
func Age(s float64, p Planet) float64 {
	var spaceAge float64

	switch p {
	case "Earth":
		spaceAge = s / Earth
	case "Mercury":
		spaceAge = s / (Earth * 0.2408467)
	case "Venus":
		spaceAge = s / (Earth * 0.61519726)
	case "Mars":
		spaceAge = s / (Earth * 1.8808158)
	case "Jupiter":
		spaceAge = s / (Earth * 11.862615)
	case "Saturn":
		spaceAge = s / (Earth * 29.447498)
	case "Uranus":
		spaceAge = s / (Earth * 84.016846)
	case "Neptune":
		spaceAge = s / (Earth * 164.79132)
	default:
		spaceAge = 0
	}

	return spaceAge
}
