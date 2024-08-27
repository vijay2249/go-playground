package arraysslicesmaps

import "fmt"

//understanding - arrays, slices, maps

func Arrays(){
	prices := [4]float32{10, 23, 99.99, 0.0}
	fmt.Println(prices)

	var books [3]string
	books = [3]string{"7 Habits of Highly Effective People"}
	fmt.Println(books)

	books[2] = "Atomic Habits"
	fmt.Println(books[0])
	fmt.Println(books)
}

/*
when we create slice - we are not creating new array or new object per-se
we are just referencing to that small data/object in the parent array or object
*/
func Slices(){
	prices := [4]float32{10, 23, 99.99, 0.0}
	fmt.Println(prices)

	//we cant use -ve index like we use in python
	featuredPrices := prices[1:3] //start with including element with index 1 and till index 3 excluding this element in index 3
	fmt.Println(featuredPrices)
	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)
	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)
	anotherSlice := featuredPrices[1:] //slice of slice
	fmt.Println(anotherSlice)
}

func DeepDiveSlices(){
	prices := [4]float32{10, 23, 99.99, 0.0}
	fmt.Println(prices)

	//we cant use -ve index like we use in python
	featuredPrices := prices[1:3] //start with including element with index 1 and till index 3 excluding this element in index 3
	fmt.Println(featuredPrices)
	featuredPrices[1] = 199.99
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices)) //cap - capacity of the slice or array

	hh := featuredPrices[:1]
	fmt.Println(hh)
	fmt.Println(len(hh), cap(hh))
	hh = hh[:3]
	fmt.Println(hh)
	fmt.Println(len(hh), cap(hh))
}

func DynamicArrays() {
	books := []string{"7 Habit", "Atomic Habits"}
	books[1] = "Alchemist"
	fmt.Println(books)
	newBooks := append(books, "Rich Dad, Poor Dad")
	fmt.Println(books)
	fmt.Println(newBooks)

}