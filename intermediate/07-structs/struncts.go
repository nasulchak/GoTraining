package main

import "fmt"

type Address struct {
	city    string
	country string
}
type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       3,
		address: Address{
			city:    "Londin",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "77568921",
			cell: "55566222",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p2.address.city = "New York"
	p2.address.country = "USA"

	p3 := Person{
		firstName: "Jane",
		age:       25,
		address: Address{
			city:    "New York",
			country: "USA",
		},
	}
	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p2.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)
	fmt.Println("Are p1 and p2 equal:", p1 == p2)
	fmt.Println("Are p3 and p2 equal:", p3 == p2)

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Println(user.username)
	p1.incrementAgeByOne()
	fmt.Println(p1.age)
}
