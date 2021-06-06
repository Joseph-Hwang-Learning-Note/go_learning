package note

import (
	"fmt"
	"strings"
)

func canIDring(age int) bool {
	switch koreanAge := age + 2; koreanAge { // can be applied at if statement. This is called variable expression
	case 10: // expression can be inserted when, in this case, there is no koreanAge at the previous line
		return false
	case 18:
		return true
	}
	return false
}

func superAdd(numbers ...int) { // ... means array
	for index, number := range numbers { // cf) for i := 0; i < len(numbers); i++
		fmt.Println(index, number)
	}
}

func lenAndUpper(name string) (length int, uppercase string) { // you can set return value like this too
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func pointerExample() {
	a := 2
	b := &a             // pointer
	fmt.Println(&a, &b) // pointer. will print the memory address of each
	fmt.Println(*b)     // see through. will print: 2
	*b = 20             // a will be changed to 20
}

func arrayExample() {
	numberArray := [5]int{1, 2, 3, 4, 5} // array: you need to specify the length
	fmt.Println(numberArray)
	numberSlice := []int{1, 2, 3, 4, 5} // slice: you don't have to specify the length
	numberSlice = append(numberSlice, 6)
	fmt.Println(numberSlice)
}

func mapExample() {
	joseph := map[string]string{
		"name": "joseph",
		"age":  "20",
	}
	fmt.Println(joseph)
	for key, value := range joseph {
		fmt.Println(key, value)
	}
}

func structExample() {
	type person struct {
		name     string
		age      int
		favorite []string
	}

	favFood := []string{"steak", "pasta"}
	joseph := person{"joseph", 20, favFood}
	// joseph := person{name: "joseph",age: 20,favorite: favFood} => Alternative way
	fmt.Println(joseph)
}
