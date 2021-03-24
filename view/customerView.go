package main

import (
	"customerManagement/service"
	"fmt"
)

type customerView struct{
	// the user input
		key string
	// whether show the menu
		exit bool
		customerService *service.CustomerService
}

//show customer list
func (this *customerView) showList() {
	customers := this.customerService.List()
	fmt.Println("------------ Customer List -----------")
	fmt.Println("ID\tname\tgender\tage\tphone\temail")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetCustomerDetails())
	}
	fmt.Println("------------ The end of List")
}
//show menu
func(this *customerView) mainMenu(){
	for{
		fmt.Println("------- Customer Info Management System -------")
		fmt.Println("------- 1. Add Customers")
		fmt.Println("------- 2. Delete Customers")
		fmt.Println("------- 3. Delete Customers")
		fmt.Println("------- 4. Check Customer list")
		fmt.Println("------- 5. Exit")
		fmt.Println("------- Please choose 1-5")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("1")
		case "2":
			fmt.Println("2")
		case "3":
			fmt.Println("3")
		case "4":
			this.showList()
		case "5":
			this.exit = true
		default:
			fmt.Println("Please enter the number from 1 to 5")
		}
		if this.exit {
			fmt.Println("Thank you for using the tool")
			break
		}
	}
	
}
func main(){ 
	//init the object
	customerView := customerView{
		key : "",
		exit: false,
	}
	customerView.customerService = service.NewCustomerService()
	//call the function
	customerView.mainMenu();
}


