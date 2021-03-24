package service

import "customerManagement/model"

// crud customer data
type CustomerService struct {
	customers []model.Customer	
	// the number of customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum= 1
	customer := model.NewCustomer(1, "Mary","Female",58,"123455","mary@email.com")
	customerService.customers = append(customerService.customers, customer);
	return customerService
}

//return customer list
func (this *CustomerService) GetCustomerList() []model.Customer {
	return this.customers
}

func(this *CustomerService) AddCustomer(customer Customer) bool {
	this.customerNum += 1
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true;
} 