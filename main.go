package main

import (
	"fmt"
	"go-rename-example/greeting"
)

func main() {
	greet := greeting.NewGreeting()
	h := greet.SayHello()
	fmt.Println(fmt.Sprint(h.Name, h.LastName))
}
