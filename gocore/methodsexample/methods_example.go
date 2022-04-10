package methodsexample

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func SimpleMethodsTest() {
	var k kelvin = 249.0
	var f fahrenheit = 23
	var c celsius
	c = kelvinToCelsius(k)
	c = k.celsius()
	c = f.celsius()
	fmt.Print(k, "* K is ", c, "* C")
}
