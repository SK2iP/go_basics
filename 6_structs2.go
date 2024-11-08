package main

import "fmt"

type messageToSend struct {
	message  string
	sender user
}

type user struct {
	name string
	phoneNumber int
}

func test(msg messageToSend) {
	fmt.Printf("%s is sent to %d phone number by %s\n", msg.message, msg.sender.phoneNumber, msg.sender.name)
}

func main() {
	msg := messageToSend{"Hello World", user{"John Doe", 1234567890}}
	test(msg)
	new_msg := messageToSend{message: "Hello World", sender: user{name: "John Doe", phoneNumber: 1234567890}}
	test(new_msg)
	front_msg := messageToSend{}
	front_msg.message = "Hello World"
	front_msg.sender.name = "John Doe"
	front_msg.sender.phoneNumber = 1234567890
	test(front_msg)
}
