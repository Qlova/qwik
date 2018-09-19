package main

import (
	"os"
	"fmt"
	"path"

	"github.com/qlova/script/compiler"
	"github.com/qlova/script/language/go"
	"github.com/qlova/i/syntax"
)

func build(i string) {
	
	file, err := os.Open(i)
	if err != nil {
		fmt.Println(err)
		return
	}

	var compiler = compiler.New()
	var language = Go.Language()

	compiler.SetSyntax(ilang.Syntax)
	
	compiler.AddInput(file)
	var program = compiler.Compile()
	
	if os.Args[1] == "go" {
		source, err := program.Source(language)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(source)
		return
	}
	
	//Interpreter.
	if os.Args[1] == "run" {
		err = program.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	
	var OutputFilePath string = "./main.go"
	if os.Args[1] == "build" {
		defer func() {
			err = os.Remove(OutputFilePath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	
		err = program.WriteToFile(OutputFilePath, language)
		if err != nil {
			fmt.Println(err)
			return
		}
			
		output, err := language.Build(OutputFilePath).CombinedOutput()
		if err != nil {
			os.Stdout.Write(output)
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: qwik [build/run]")
		return
	}
	
	switch os.Args[1] {
		case "build", "run", "go":
			
			directory, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			
			name := path.Base(directory)
			
			if len(os.Args) < 3 {
				build(directory+"/"+name+".i")
			} else {
				build(os.Args[2])
			}
			
		default:
			fmt.Println("Usage: qwik [build/run]")
	}
}
