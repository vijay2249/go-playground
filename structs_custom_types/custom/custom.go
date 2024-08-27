package custom

import "fmt"

type Number int

func (n Number) log() {
	fmt.Println(n)
}

func Start(){
	var age Number = 30
	age.log()
}