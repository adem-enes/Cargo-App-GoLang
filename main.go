package main

import (
	"fmt"
	"lesson1/structure"
	"os"
	"os/exec"
)

// Statuses
// gönderiliyor packaging
// teslim edildi delivered
// geri iade returned
//

var customer1 = structure.NewCustomer(1, "adem", "polat", 555, "Burası")
var customer2 = structure.NewCustomer(2, "sevket", "yılmaz", 555, "Şurası")

//If we want to update the addresses of customers update anything about customer we have to use pointers because
// the customer update also has to change in the orders..
var customers = map[int]*structure.Customer{
	customer1.IdNumber: &customer1,
	customer2.IdNumber: &customer2,
}

var order1 = structure.NewOrder("gönderiliyor", customers[1], customers[2])
var order2 = structure.NewOrder("teslim edildi", customers[2], customers[1])

//We might be use the pointer for key values..
var orders = map[int]structure.Order{
	order1.Id: order1,
	order2.Id: order2,
}

func main() {
	menu()
}

func menu() {
	clearScreen()
	fmt.Println("\t\t\t..::Welcome Cargo App::..")
	fmt.Println("[1] - See Orders \n[2] - See Customers \n[3] - Create New Order")
	fmt.Println("[4] - Change Order Status \n[5] - Change Customer Address\n[6] - Order Check")
	fmt.Println("[7] - Exit")
	fmt.Print("Your Choise: ")
	var choise int
	fmt.Scan(&choise)

	switch choise {
	case 1: //See Orders
		clearScreen()
		seeOrders()
		returnMenu()
	case 2: //See Customers
		clearScreen()
		seeCustomers()
		returnMenu()
	case 3: //Create New Order
		clearScreen()
		createNewOrder()
		returnMenu()
	case 4: //Change Order Status
		clearScreen()
		orderStatusUpdate()
		returnMenu()
	case 5: // Change Customer Address
		clearScreen()
		customerAddressUpdate()
		returnMenu()
	case 6: //Order Check
		clearScreen()
		checkOrder()
		returnMenu()
	case 7: //Exit
		clearScreen()
	default:
		clearScreen()
		fmt.Println("Wrong Choise")
		returnMenu()
	}
}
func returnMenu() {
	fmt.Println("\n---------------------------------------------------------------------")

	fmt.Println("\nTo Return Menu Please Press 'Y'")
	var choise string
	fmt.Scan(&choise)
	if choise == "Y" || choise == "y" {
		clearScreen()
		menu()
	}
}
func clearScreen() {
	//To clear console in windows..
	// cmd := exec.Command("cmd", "/c", "cls")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	// To clear console in Mac or Linux
	// fmt.Println("\033[2J")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func seeOrders() {
	fmt.Println("\nOrders: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Order ID \t\t Order Status \t\t Order Receiver \t\t\t Order Sender")
	fmt.Println("------------------- \t ---------------\t----------------------------------\t---------------------------")

	for k, v := range orders {
		fmt.Print(k)
		fmt.Println("\t", v.Status, "  \t", v.Receiver, "\t\t ", v.Sender)
	}
}

func seeCustomers() {
	fmt.Println("\nCustomers: ")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Customer ID \t\t Name \t\t\t LastName \t\t Phone Number \t\t Address")
	fmt.Println("--------------- \t ---------------\t-----------------\t------------------\t--------------")

	for k, v := range customers {
		fmt.Print(k, "\t\t")
		fmt.Println("\t", v.Name, "  \t\t", v.LastName, "  \t\t ", v.PhoneNumber, "\t\t\t ", v.Address)
	}
}

func createNewOrder() {
	fmt.Println("Sender Values:")
	fmt.Println("---------------------------------------------------------------------")
	var newSender structure.Customer

	fmt.Print("Sender Name: |\t")
	fmt.Scan(&newSender.Name)

	//Add new Customer
	customers[newSender.IdNumber] = &newSender
}

func checkOrder() {
	fmt.Println("\nOrder Statuses: ")
	fmt.Println("---------------------------------------------------------------------")
	for k, v := range orders {
		fmt.Print("Order ID: ", k)
		fmt.Println("\tOrder Status: ", v.Status)
	}

	// var inquiry string
	// fmt.Print("\nOrder Status Filter: ")
	// fmt.Scan(&inquiry)

	// fmt.Println("Order Status Filter ")
	// fmt.Println("---------------------------------------------------------------------")

	// for k, v := range orders {
	// 	if v.Status == inquiry {
	// 		fmt.Println(k)
	// 	}
	// }

	var idInquiry int
	fmt.Print("\nPlease Enter Order Id You Want To Control: ")
	fmt.Scan(&idInquiry)

	fmt.Println("Order Status")
	fmt.Println("---------------------------------------------------------------------")

	for k, v := range orders {
		if k == idInquiry {
			fmt.Println("Order ID: ", k)
			fmt.Println("Order Status: ", v.Status)
			fmt.Println("--------------------")
			fmt.Println("Order Sender: ")
			fmt.Println("Sender Id: ", v.Sender.IdNumber)
			fmt.Println("Sender Name LastName: ", v.Sender.Name, " ", v.Sender.LastName)
			fmt.Println("Sender Phone Number: ", v.Sender.PhoneNumber)
			fmt.Println("Sender Address: ", v.Sender.Address)
			fmt.Println("--------------------")
			fmt.Println("Order Receiver: ")
			fmt.Println("Receiver Id: ", v.Receiver.IdNumber)
			fmt.Println("Receiver Name LastName: ", v.Receiver.Name, " ", v.Receiver.LastName)
			fmt.Println("Receiver Phone Number: ", v.Receiver.PhoneNumber)
			fmt.Println("Receiver Address: ", v.Receiver.Address)
		}
	}
}

func customerAddressUpdate() {
	fmt.Println("__Customer Address Update__")
	seeCustomers()

	fmt.Println("\nCustomer Address Update")
	fmt.Println("---------------------------------------------------------------------")

	var customerId int
	var newAddress string
	fmt.Print("\nPlease Enter the customers id: \t")
	fmt.Scan(&customerId)

	fmt.Print("Please Enter the new Address: \t")
	fmt.Scan(&newAddress)

	newCustomer := customers[customerId]
	// customers[customerId].SetCustomerAddress(newAddress)
	// customers[customerId].Address = newAddress
	newCustomer.SetCustomerAddress(newAddress)
	delete(customers, customerId)
	customers[newCustomer.IdNumber] = newCustomer
	// fmt.Println("Customers new List")
	// fmt.Println("---------------------------------------------------------------------")
	// for k, v := range customers {
	// 	fmt.Println(k, ": \t", v)
	// }
	fmt.Print("\nCustomer Address of ", customerId, " is cahnged to '", newAddress, "' successfully..")
}

func orderStatusUpdate() {
	fmt.Println("__OrderStatus Update__")
	seeOrders()

	fmt.Println("\nOrder Status Update")
	fmt.Println("---------------------------------------------------------------------")

	var orderId int
	var statusChoise int
	fmt.Print("\nPlease Enter the Order Id: ")
	fmt.Scan(&orderId)

	fmt.Println("\t..Statuses..")
	fmt.Println("[1] - Gönderiliyor\n[2] - Teslim Edildi\n[3] - İade")
	fmt.Print("Please Choose New Status: ")
	fmt.Scan(&statusChoise)
	var newStatus string
	if statusChoise == 1 {
		newStatus = "Gönderiliyor"
	} else if statusChoise == 2 {
		newStatus = "Teslim Edildi"
	} else if statusChoise == 3 {
		newStatus = "İade"
	} else {
		fmt.Println("Wrong Choise.. Default Status is set to default (Status-1)")
		newStatus = "Gönderiliyor"
	}

	newOrder := orders[orderId]
	// customers[customerId].SetCustomerAddress(newAddress)
	// customers[customerId].Address = newAddress
	newOrder.SetOrderStatus(newStatus)
	delete(customers, orderId)
	orders[newOrder.Id] = newOrder

	fmt.Print("\nOrder Status of ", orderId, " is cahnged to '", newStatus, "' successfully..")
}
