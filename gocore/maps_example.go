package gocore

import "fmt"

func MapsBasicExample() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	temperature["Earth"] = 16
	temperature["Venera"] = 464 // add pair
	fmt.Println(temp)
	fmt.Println(temperature["Legion"]) // return 0
	fmt.Println(temperature)
	fmt.Println("----------comma ok---------")
	moon := temperature["Moon"]
	fmt.Println(moon)
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("dT %v", moon)
	} else {
		fmt.Println("Where is moon?")
	}
}

func MapsInMemoryExample() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}

func MapsMake() {
	temperature := make(map[float64]int, 8) // 8 as capacity
	fmt.Println(temperature)
}

func MapsFreqExample() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%.2f count %d\n", t, num)
	}
}
