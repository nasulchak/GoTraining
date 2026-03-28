package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromreader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}

	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing from reader.", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing file")
	}
}

func bufferExample() {
	var buf bytes.Buffer // stack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multipleReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("Wrold!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe."))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file:", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to file:", err)
	}

	// Type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error writing to file:", err)
	// }
}

func main() {
	fmt.Println("=== Read from Reader ===")
	readFromreader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer!")
	fmt.Println(writer)

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader ===")
	multipleReaderExample()

	fmt.Println("=== Pipe ===")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "Hello File")

	resource := &MyResoutce{name: "TestResource"}
	closeResource(resource)
}

type MyResoutce struct {
	name string
}

func (m MyResoutce) Close() error {
	fmt.Println("Clossing resource:", m.name)
	return nil
}
