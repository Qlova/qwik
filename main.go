package main

import (
	"os"
	"fmt"
	"path"

	"github.com/qlova/script/compiler"
	"github.com/qlova/script/languages/go"
	"github.com/qlova/i/syntax"
	
	"io/ioutil"
)

func build(i string) {
	
	file, err := os.Open(i)
	if err != nil {
		fmt.Println(err)
		return
	}

	var compiler = compiler.New()
	var language = Go.Language()
	
	compiler.Script.SetLanguage(language)
	
	compiler.SetSyntax(ilang.Syntax)
	
	compiler.AddInput(file)
	compiler.Compile()
	
	if compiler.Errors {
		return
	}
	
	var OutputFilePath string = "./"+Go.DefaultFileName
	if os.Args[1] == "build" {
		defer func() {
			err = os.Remove(OutputFilePath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
	
	if os.Args[1] == "run" {

		TemporaryDirectory, err := ioutil.TempDir("", "qwik")
		if err != nil {
			fmt.Println(err)
			return
		}
		
		defer func() {
			err = os.RemoveAll(TemporaryDirectory)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
		
		OutputFilePath = TemporaryDirectory+"/"+Go.DefaultFileName
	}
		
	OutputFile, err := os.Create(OutputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	compiler.Script.WriteTo(OutputFile)
	OutputFile.Close()
	
		
	output, err := language.Build(OutputFilePath).CombinedOutput()
	if err != nil {
		os.Stdout.Write(output)
		return
	}
	
	if os.Args[1] == "run" {
		cmd := language.Run(OutputFilePath)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		
		if err != nil {
			fmt.Println(err)
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
		case "build", "run":
			build()
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
