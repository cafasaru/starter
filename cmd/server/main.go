package main

import "fmt"




func Run() error {
	fmt.Println("running service")
	return nil
}
func main() { 
	fmt.Println("starting server")

	if err := Run(); err != nil {
		fmt.Println(err.Error())
	}
} 
