package arraysslicesmaps

import "fmt"

type Product struct {
	title string
	id int
	price float32
}

func Exercise(){
	hobbies := [3]string{"Music", "Teaching", "Read"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[1])

	newSlice := hobbies[:2]
	anotherNewSlice := hobbies[0:2]
	fmt.Println(newSlice)
	fmt.Println(anotherNewSlice)

	anotherNewSlice = anotherNewSlice[1:]
	fmt.Println(anotherNewSlice)
	fmt.Println(len(anotherNewSlice), cap(anotherNewSlice))

	newHobbies := []string{"Live", "Fun"}
	newHobbies[1] = "Laugh"
	newHobbies = append(newHobbies, "Learn")
	fmt.Println(newHobbies)

	allHobbies := append(newHobbies, hobbies[:]...)
	fmt.Println(allHobbies)

	products := []Product{
		{title: "Tiny Habits", id: 5, price: 599.0},
		{title: "7 Habits", id: 1, price: 799.00},
	}

	fmt.Println(products)
	fmt.Println(products[1].title)
	products = append(products, Product{title: "Alchemist", id: 4, price: 249.99}, Product{title: "Think and Grow Big", id: 99, price: 139.99})
	fmt.Println(products)
}