package main

import "fmt"

type User struct {
	id        int
	firstName string
	lastName  string
	email     string
	address   Address
	Status
}

type Address struct {
	city     string
	province string
}

type Status struct {
	active bool
}

func main() {
	user1 := User{}
	user1.id = 1
	user1.firstName = "Annisa"
	user1.lastName = "Nadia"
	user1.email = "annisanadianeyla@gmail.com"
	user1.address.city = "Bekasi"
	user1.address.province = "West Java"
	user1.active = true

	user2 := User{
		id:        2,
		firstName: "John",
		lastName:  "Mayer",
		email:     "johnmayer@gmail.com",
		address: Address{
			city:     "Bandung",
			province: "West Java",
		},
		Status: Status{
			active: true,
		},
	}

	fmt.Printf("%+v\n", user1)
	fmt.Printf("id: %v\n", user1.id)
	fmt.Printf("firstName: %v\n", user1.firstName)
	fmt.Printf("lastName: %v\n", user1.lastName)
	fmt.Printf("email: %v\n", user1.email)
	fmt.Printf("city: %v\n", user1.address.city)
	fmt.Printf("province: %v\n", user1.address.province)
	fmt.Printf("province: %v\n\n", user1.active)

	fmt.Printf("%+v\n", user2)
	fmt.Printf("id: %v\n", user2.id)
	fmt.Printf("firstName: %v\n", user2.firstName)
	fmt.Printf("lastName: %v\n", user2.lastName)
	fmt.Printf("email: %v\n", user2.email)
	fmt.Printf("city: %v\n", user2.address.city)
	fmt.Printf("province: %v\n", user2.address.province)
	fmt.Printf("province: %v\n\n", user2.active)
}
