package main

import (
	"fmt"
	greetings "github.com/intliang/go_greetings"
)

func main() {
	msg, err := greetings.Hello("intliang")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
}
