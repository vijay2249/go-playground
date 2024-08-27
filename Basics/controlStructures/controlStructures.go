package controlstructures

import (
	"fmt"

	"wedonttrack.org/learn/go/basics/Basics/files"
)

func outputOptions() {
	fmt.Println("Welcome to fSociety Bank")
	
	fmt.Println("Waht do you want to do: ")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func Controlstructures(){
	accountBalance, err := files.GetBalanceToFile()
	if err != nil {
		fmt.Println("error")
		fmt.Print(err)
		fmt.Println("====")
		panic(err)
	}

	outputOptions()


	var choice int

	for {
		fmt.Print("your choice: ")
		fmt.Scan(&choice)

		checkBalance := choice == 1

		//switch conditions in GO
		// switch choice {
		// case 1:
		// 	fmt.Printf("Account Balane: %v\n", accountBalance)
		// case 2:
		// 	fmt.Print("Deposit amount: ")
		// 	var deposit float64
		// 	fmt.Scan(&deposit)
		// 	if deposit <= 0 {
		// 		fmt.Println("Value must be >=0 ")
		// 		continue
		// 	}
		// 	accountBalance += deposit
		// 	fmt.Println("New account balance: ", accountBalance)
		// 	files.WriteBalanceToFile(accountBalance)
		// case 3:
		// 	fmt.Print("Withdraw amount: ")
		// 	var withdraw float64
		// 	fmt.Scan(&withdraw)
		// 	if withdraw <= 0 {
		// 		fmt.Println("Value must be >=0 ")
		// 		continue
		// 	}
		// 	if withdraw > accountBalance {
		// 		fmt.Println("Invalid amount, cant withdraw amount more than account balance")
		// 		continue
		// 	}
		// 	accountBalance -= withdraw
		// 	fmt.Println("New account balance: ", accountBalance)
		// 	files.WriteBalanceToFile(accountBalance)
		// default:
		// 	fmt.Println("Bye Bye")
		// 	return
		// }


		//if-else blocks in GO
		if checkBalance {
			fmt.Printf("Account Balane: %v\n", accountBalance)
		} else if choice == 2 {
			fmt.Print("Deposit amount: ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Value must be >=0 ")
				continue
			}
			accountBalance += deposit
			fmt.Println("New account balance: ", accountBalance)
			files.WriteBalanceToFile(accountBalance)
		} else if (choice == 3){
			fmt.Print("Withdraw amount: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Value must be >=0 ")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Invalid amount, cant withdraw amount more than account balance")
				continue
			}
			accountBalance -= withdraw
			fmt.Println("New account balance: ", accountBalance)
			files.WriteBalanceToFile(accountBalance)
		} else {
			fmt.Println("Bye Bye")
			break
		}
	}

	fmt.Println("Bye Bye from fSociety Bank")	
}

func infiniteLoop(){
	//infinite loops
	for {
		fmt.Println("This is infinite loop in GO")
	}
}
	
	
func Loops(){
	//for loops
	for i :=0; i<3; i++ {
		fmt.Println("For Loops in GO")
	}
}