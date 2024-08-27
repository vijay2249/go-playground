package generics

import "fmt"


func Main() {
	result := addSingleType(1,2)
	anotherResult := addAnyType(2,3)
	genericResult := addGenericTypes(3,5)

	fmt.Println(result, anotherResult, genericResult)
}

func addSingleType(a int, b int) int { return a + b }

func addGenericTypes[T int|float32|float64|string] (a, b T) T { return a + b }

func addAnyType(a, b interface{}) interface{}{
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt { return aInt + bInt }

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat { return aFloat + bFloat }
	return nil
}