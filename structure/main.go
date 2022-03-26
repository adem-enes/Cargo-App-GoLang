package structure

import (
	"math/rand"
)

type Order struct {
	Id       int
	Status   string
	Receiver *Customer
	Sender   *Customer
}

type Customer struct {
	IdNumber    int
	Name        string
	LastName    string
	PhoneNumber int
	Address     string
}

func NewCustomer(IdNumber int, Name, LastName string, PhoneNumber int, Address string) Customer {
	return Customer{
		IdNumber:    IdNumber,
		Name:        Name,
		LastName:    LastName,
		PhoneNumber: PhoneNumber,
		Address:     Address,
	}
}

func NewOrder(status string, reciver, sender *Customer) Order {
	return Order{
		Id:       rand.Int(),
		Status:   status,
		Receiver: reciver,
		Sender:   sender,
	}
}

// Order ID
func (order *Order) SetOrderStatus(status string) {
	order.Status = status
}

//Customer Adress Update
func (customer *Customer) SetCustomerAddress(address string) {
	customer.Address = address
}
