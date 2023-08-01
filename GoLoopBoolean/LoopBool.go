package main 

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	for i := 0; i< 5; i++ {
		fmt.Println("The value of i is: ", i)
	}

	names := []string{"Yoshi", "Mario", "Luigi", "Bowser"}

	for i:= 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("The position at index %v is %v \n", index, value)
	}

	// Boolean 

	age := 56

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	for index, val := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
	}
}