package functions

import "fmt"

type transformFnType func(int) int

func Main(){
	numbers := []int{22,23,24}
	
	dNumbers := factorNumbers(&numbers, double)
	tNumbers := factorNumbers(&numbers, triple)
	fmt.Println(dNumbers)
	fmt.Println(tNumbers)
	
	anotherNum := []int{1,2,3}
	anotherTransformerFunction := getTransformer(&anotherNum)
	transformerFunction := getTransformer(&numbers)

	anotherTransformedNum := factorNumbers(&anotherNum, anotherTransformerFunction)
	transformedNum := factorNumbers(&anotherNum, transformerFunction)

	fmt.Println(transformedNum)
	fmt.Println(anotherTransformedNum)

}

func factorNumbers(num *[]int, transform transformFnType) []int{
	dNumbers := []int{}
	for _, val := range *num {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformer(num *[]int) transformFnType {
	if (*num)[0] == 1 { 
		return double
	} else {
		return triple
	}
}

func double(num int) int {
	return num*2
}

func triple(num int) int {
	return num*3
}

func AnonymousFunctions() {
	numbers := []int{1,2,3}
	//                                  > ANONYMOUS FUNCTION               <
	transfomed := factorNumbers(&numbers, func (num int) int {return num*2})
	fmt.Println(transfomed)
}


//clousers
func Clousers() {
	numbers := []int{2,3,4}
	double := createTransformer(2)
	triple := createTransformer(3)
	
	doubled := factorNumbers(&numbers, double)
	tripled := factorNumbers(&numbers, triple)
	
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num*factor
	}
}


//Recusions
func Recursions() {
	fmt.Println(factorialLoop(6))
	fmt.Println(recursiveFactorial(4))
}

func recursiveFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

func factorialLoop(n int) int {
	result := 1
	for i:=1; i<=n; i++ {
		result = result*i
	}
	return result
}