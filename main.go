package main

import (
	"fmt"
	"goLangProject/httpGetPost"
)

func main() {
	fmt.Println("main msg")
	//smtp.EmailSendMain()
	//httpGetPost.PostEx()

	//httpGetPost.JsonPost(struct {
	//	Name string
	//	Age  int
	//}{Name: "Alex", Age: 20})

	httpGetPost.HttpClient()
}
