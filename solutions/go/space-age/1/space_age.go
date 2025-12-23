package space

type Planet string

const seconds_in_year int = 31557600

var planetsOrbitalPeriod map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	orbitalPeriod, exists := planetsOrbitalPeriod[planet]
	if !exists {
		return -1
	}

	return seconds / float64(seconds_in_year) / orbitalPeriod
}
