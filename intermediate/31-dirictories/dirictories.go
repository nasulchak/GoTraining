package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// checkError(os.Mkdir("subdir", 0755))
	// defer os.RemoveAll("subdir")

	// os.WriteFile("subdir/file", []byte(""), 0755)

	os.MkdirAll("subdir/parent/child", 0755)
	os.MkdirAll("subdir/parent/child1", 0755)
	os.MkdirAll("subdir/parent/child2", 0755)
	os.MkdirAll("subdir/parent/child3", 0755)

	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	fmt.Println("Readind subdir/parent:")
	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	fmt.Println("Readind subdir/parent/child:")
	checkError(os.Chdir("subdir/parent/child"))
	result, err = os.ReadDir(".")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry)
	}

	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// filepath.Walk and filepatk.WalkDir
	pathfile := "subdir"
	fmt.Println("Walking Directory")
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Println("->", path)
		return nil
	})

	checkError(err)

	checkError(os.RemoveAll("./subdir"))
}
