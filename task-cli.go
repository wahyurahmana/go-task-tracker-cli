package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//GET ARGUMENT COMMAND
	input := os.Args

	//MAPPING COMMAND
	if strings.ToLower(input[1]) == "add" {
		fmt.Println("ini adalah proses add")
	} else if strings.ToLower(input[1]) == "update" {
		fmt.Println("ini adalah proses update")
	} else if strings.ToLower(input[1]) == "delete" {
		fmt.Println("ini adalah proses delete")
	} else if strings.ToLower(input[1]) == "mark-in-progress" {
		fmt.Println("ini adalah proses mark-in-progress")
	} else if strings.ToLower(input[1]) == "mark-done" {
		fmt.Println("ini adalah proses mark-done")
	} else if strings.ToLower(input[1]) == "list" {
		if len(input[1:]) > 1 {
			if strings.ToLower(input[2]) == "done" {
				fmt.Println("ini adalah proses done")
			} else if strings.ToLower(input[2]) == "todo" {
				fmt.Println("ini adalah proses todo")
			} else if strings.ToLower(input[2]) == "in-progress" {
				fmt.Println("ini adalah proses in-progress")
			} else {
				fmt.Println("Please Input Your Command")
			}
		} else {
			fmt.Println("ini adalah proses list")
		}
	} else {
		fmt.Println("Please Input Your Command")
	}
}
