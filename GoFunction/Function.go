package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye!! %v \n", n)
}

func cycleName(n []string, f func(string )) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi *r *r 
}
func main() {
	sayGreeting("Dhyan Shah")
	sayGreeting("Madhav Sampat")
	sayBye("Dhayn Shah")

	cycleName([]string{"Cloud", "Tiffer", "Barret"}, sayGreeting)
	cycleName([]string{"Cloud", "Tiffer", "Barret"}, sayBye)

	a1 := circleArea(12.6)
	a2 := circleArea(15.4)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
