package main

import "fmt"

func main(){

// Slice declaration and initialization
var val []int
var values = []int{}
var nums = []int{1,2,3}
numbers := []int{}
countingnums := []int{1,2,3,4,5}

fmt.Printf("Vals: %v\n", val)
fmt.Printf("Values: %v\n", values)
fmt.Printf("Nums: %v\n", nums)
fmt.Printf("Numbers: %v\n", numbers)
fmt.Printf("Counting Numbers: %v\n", countingnums)


fmt.Println(val == nil) // => true
fmt.Println(values == nil) // => false


// Slice declaration and initialization using make function
names := make([]string, 4,5)

// Slice length and capacity

fmt.Printf("The length of the names slice is %v and the capacity is %v\n", len(names), cap(names))

// Inserting elements in slice

fmt.Printf("Names: %v\n", names)
names[1] = "John"

fmt.Printf("Names: %v\n", names)

//Append
slice := []int{1, 2, 3}

slice = append(slice, 4,5)

fmt.Printf("Slice values: %v\n", slice)

//Replaces the value
slice[2] = 6

fmt.Printf("Slice values: %v\n", slice)

//What is I want to add without replacing existing value?
slicenums := []int {1,2,3,4,5}
index := 2
value := 12

fmt.Printf("Slice Nums values: %v\n", slicenums)
slicenums = append(slicenums[:index],append([]int {value},slicenums[index:]...)...)
fmt.Printf("Slice Nums values: %v\n", slicenums)

//Subslice
subslice := slice[1:3]

subslice2 := slice[1:4:4]

fmt.Printf("Slice length: %v\n", len(slice))
fmt.Printf("Slice capacity: %v\n", cap(slice))

fmt.Printf("Subslice values: %v\n", subslice)

fmt.Printf("Subslice2 values: %v\n", subslice2)

// Accessing elements of slice

fmt.Printf("Slice value at index 1: %v\n", slice[1])

// Deleting elements from slice

fmt.Printf("Slice Nums values: %v\n", slicenums)
slicenums = append(slicenums[:index],slicenums[index+1:]...)
fmt.Printf("Slice Nums values: %v\n", slicenums)

// Copying elements of slice

namez := []string {"John", "Doe"}
copiedNamez := namez

fmt.Printf("List of names: %v\n", namez)
fmt.Printf("List of copied names: %v\n", copiedNamez)

namez[0] = "Jane"

fmt.Printf("List of names: %v\n", namez)
fmt.Printf("List of copied names: %v\n", copiedNamez)

// Deep copy of slice

cartoons := []string {"Johnny Bravo", "Mickey Mouse"}
newCartoons := []string{}
newCartoons = append(newCartoons,cartoons...)
//Note order

fmt.Printf("List of cartoons: %v\n", cartoons)
fmt.Printf("List of new cartoons: %v\n", newCartoons)

newCartoons[1] = "Cinderella"
fmt.Printf("List of cartoons: %v\n", cartoons)
fmt.Printf("List of new cartoons: %v\n", newCartoons)


}