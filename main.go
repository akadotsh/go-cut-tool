package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)


type Commands struct{
	
}


func (c *Commands)Run(command string){
	if(command == "-f2"){

	}
}


func main (){

	var command string
	var delimiter string
	var fields string

	// Parse flags
	flag.StringVar(&command,"cut", "", "Command to run")
	flag.StringVar(&delimiter, "d", "\t", "Delimiter character")
	flag.StringVar(&fields, "f", "-f1", "Field numbers to extract")
	flag.Parse()

	
	args := flag.Args()
	
	
	if flag.NArg() < 1 || flag.Arg(0) != "cut" {
		fmt.Println("Invalid commands. Please specify a valid command.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	for _,c := range args{
		if strings.HasPrefix(c,"-d"){
			delimiter = c[2:]
			break
		}
	}


	fmt.Println("delimiter",delimiter)

	fmt.Println("fields",fields)

	filename := args[len(args)-1]

	file, err := os.Open(filename)

	if err!=nil{
			log.Fatal("Error: ",err)
	}

	defer file.Close()


	scanner :=bufio.NewScanner(file);
	



	if fields == "-f1" {
		fmt.Println("filename",filename)

		for	scanner.Scan(){
		   parts:= strings.Split(scanner.Text(),delimiter);
		   fmt.Println(parts[0])
		}
	}else if fields == "-f2"{
		for	scanner.Scan(){
		   parts:= strings.Split( scanner.Text(),"\t");
		   fmt.Println(parts[1])
		}

	}
	
}

 
	
	 
	