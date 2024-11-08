package main

import "fmt"

//create an example that teach me error using the fmt.Errorf function and error interface no create the error interface
type error interface {
	Error() string
}

type MessagetoClient struct {
	Msg string
}

func messagecheck(m MessagetoClient) (string , error) {
	if m.Msg == "" {
		return "", fmt.Errorf("Error: Message is empty")
	}
	return m.Msg, nil
}

func main() {
	m := MessagetoClient{Msg: ""}
	new_msg , err := messagecheck(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(new_msg)
}