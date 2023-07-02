package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	dirPath := "./schema"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if !file.IsDir() && fileName[len(fileName)-3:] == ".go" {
			generate(dirPath + "/" + fileName)
		}
	}
}

func generate(filePath string) {
	//Create a FileSet to work with
	fset := token.NewFileSet()
	//Parse the file and create an AST
	file, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		panic(err)
	}

	g := generator{}

	f := g.generate(file)

	fmt.Printf("%#v", f)
}
