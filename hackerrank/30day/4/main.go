package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	p.age = initialAge
	if initialAge < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13 {
		fmt.Println("You are young.")
	} else {
		if p.age >= 13 && p.age < 18 {
			fmt.Println("You are a teenager.")
		} else {
			fmt.Println("You are old.")
		}
	}
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age = p.age + 1
	return p
}

func main() {
	var num, age int
	// number of test cases
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
