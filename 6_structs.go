package main

import "fmt"

type messageToSend struct {
	message string 
	phoneNumber int
}

func test (msg messageToSend){
	fmt.Printf("%s is sent to %d phone number\n",msg.message,msg.phoneNumber)
}

func main(){
	msg := messageToSend{"Hello World",1234567890}
	test(msg)
	new_msg := messageToSend{message: "Hello World",phoneNumber: 1234567890}
	test(new_msg)
}