package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	// person := Person{Name: "John", Age: 30, City: "London", Email: "email@example.com"}
	person := Person{Name: "John", Address: Address{City: "London", State: "CK"}, Email: "email@example.com"}

	xmlData, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatal("Error Marshalling data into XML:", err)
	}

	fmt.Println(string(xmlData))

	// xmlRaw := `<person><name>Jane</name><age>25</age></person>`
	xmlRaw := `<person><name>Jane</name><age>25</age><address><city>Oakland</city><state>CA</state></address></person>`

	var personXML Person
	err = xml.Unmarshal([]byte(xmlRaw), &personXML)
	if err != nil {
		log.Fatal("Error Unmarshalling XML:", err)
	}

	fmt.Println(personXML)

	book := Book{
		ISBN:       "456-4-15-68",
		Title:      "Go Bootcamp",
		Author:     "Ashish",
		Pseudo:     "Pseudo",
		PseudoAttr: "PseudoAttr",
	}

	bookXML, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatal("Error Marshaling data into XML", err)
	}
	fmt.Println("Book XML:", string(bookXML))

}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoAttr,attr"`
}
