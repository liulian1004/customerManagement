package model

import "fmt"

//customer struct
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//factory pattern

func NewCustomer(id int, name string, gender string,
age int, phone string, email string ) Customer {
	return Customer{
		Id : id,
		Name :name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

func NewCustomerWithoutId(name string, gender string,
	age int, phone string, email string ) Customer {
		return Customer{
			Name :name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
		}
	}

func (this Customer) GetCustomerDetails() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}

