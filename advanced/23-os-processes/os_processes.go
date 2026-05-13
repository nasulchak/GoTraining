package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main(){
	pr, pw := io.Pipe()
	
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func(){
		defer pw.Close()
		pw.Write([]byte("food is good\nbar\nbaz"))
	}()
	
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Output: %s", output)

	// 1.
	// cmd := exec.Command("ls", "/nodaosd")
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Printf("%s", output)

	// 2.
	// cmd := exec.Command("sleep", "5")

	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error starting command:", err)
	// 	return
	// }
	
	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// 	if err != nil {
	// 	fmt.Println("Error killing process:", err)
	// 	return
	// }

	// fmt.Println("Process killed")

	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("Error waiting", err)
	// 	return
	// }

	// fmt.Println("Process is complete")

	// 3.
	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = strings.NewReader("food is good\nbar\baz")

	// output, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Output: %s", output)
	
	// 4.
	// cmd := exec.Command("echo", "Hello World!")
	// output, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Output: %s", output)
}