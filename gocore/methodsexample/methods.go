package methodsexample

import "fmt"

type report struct {
	sol
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

type sol int

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func OopMethodTest() {
	bradbury := location{-4.3, 124.3}
	t := temperature{high: 2.0, low: 13.9}
	report := report{sol: 15, temperature: t, location: bradbury}
	report.average()
	fmt.Printf("%+v\n", report)
	fmt.Println(report.days(1446))
	fmt.Println(report.sol.days(1446))
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}
