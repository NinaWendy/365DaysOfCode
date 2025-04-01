package main

import "fmt"


func changeFirst(s []int){
	s[0]= 10
}
func main(){

	// Array declaration with var keyword
	var numbers = [5]int{}
	var names = [3]string{}
	fmt.Printf("Numbers: %v \n",numbers)
	fmt.Printf("Names: %v\n", names)

	// Array initialization
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	names[0] ="John"
	names[1] ="Doe"
	names[2] ="JD"

	fmt.Printf("Numbers: %v \n",numbers)
	fmt.Printf("Names: %v\n", names)

	// Array declaration and initialization with values

	var cars = [3]string{"VW", "BMW","Audi"}
	fmt.Printf("Cars: %v\n", cars)

	// Array declaration and initialization Shorthand
	bags := [3]string{"Gucci", "Fenti","Prada"}
	fmt.Printf("Bags: %v\n", bags)

	watches := [...]string{"casio","rolex"}
	fmt.Printf("Watches: %v\n", watches)

	//Array length
	fmt.Printf("Watches array length %d\n",len(watches))
	fmt.Printf("Bags array length %d\n",len(bags))

	//Array capacity
	fmt.Printf("Watches array capacity %d\n",cap(watches))
	fmt.Printf("Bags array capacity %d\n",cap(bags))

	// Array of structs
	type person struct {
		name string
		age int
	}

	persons := [2]person{ {name: "JohnDoe", age: 20},{name: "JaneDoe", age: 30}}

	fmt.Printf("List of persons %v\n", persons)

	// Multidimensional array
	//[row][columns]
	rubics := [5][6] int{
		{1,2,3,4,5,6},
		{1,2,3,4,5,6},
		{1,2,3,4,5,6},
		{1,2,3,4,5,6},
		{1,2,3,4,5,6}}

	fmt.Printf("Rubics cube %v\n",rubics)
	

	//How to get row and columns

	rows := len(rubics)       // Number of rows (external array)
    cols := len(rubics[0])    // Number of columns (assuming all rows have the same length)
	// returns the number of elements in the first row.

    fmt.Printf("Rows: %d, Columns: %d\n", rows, cols)

	//Access array elements
	//Time complexity of O(1) since we know the elements index

	fmt.Printf("Element at index 2 of the cars array: %v\n", cars[2])

	// Update array elements

	cars[2]="Jeep"
	fmt.Printf("Element at index 2 of the cars array after update: %v\n", cars[2])

	// Array iteration with index

	for i := 0; i < len(bags); i++ {
		fmt.Printf("Bag %v is %v\n",i + 1, bags[i])
	}

	for i := 0; i < len(rubics); i++ {
		for j := 0 ; j < len(rubics[0]); j++ {
			fmt.Printf("Number at %v %v is %v\n", i , j , rubics[i][j])
		}
	}

	// Array iteration with range
    
	for index, value := range bags{
		fmt.Printf("Value at index %v is: %v\n", index, value)
	}

	for _, value := range bags{
		fmt.Println(value)
	}
	//Copy array

	newbags := bags

	// fmt.Printf("New bags:%v\n",newbags)

	newbags[1] = "Louis Vuitton"

	fmt.Printf("Old bags:%v\n", bags)
	fmt.Printf("New bags:%v\n",newbags)

	
	//Array interface
	// Arrays cannot have different data types

    combinedarray := [3] interface{} {"Nina", 4 , 32.75}

	fmt.Printf("This is a list of arrays with different types: %v\n", combinedarray)

	var x [] int = [] int{1,2,3}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)

}