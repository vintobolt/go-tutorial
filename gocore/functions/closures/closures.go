package closures

import "fmt"

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func ClosureTest() {
	sensor1 := calibrate(realSensor, 5)
	fmt.Println(sensor1())
}

func ClosureExampleValues() {
	var k kelvin = 294.0
	sensor := func() kelvin {
		return k
	}
	fmt.Println(sensor())
	k++
	fmt.Println(sensor())
}
