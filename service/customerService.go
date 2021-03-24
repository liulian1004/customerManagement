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
func (this *CustomerService) List() []model.Customer {
	return this.customers
}