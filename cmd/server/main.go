package main

import "fmt"

// responsible for startup application
func Run() error {
	fmt.Println("Starting up application")
	return nil
}

func main() {
	fmt.Println("Go comments apis")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
