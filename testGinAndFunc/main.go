package main

import (
	"github.com/gin-gonic/gin"
)

// HomePage : GET, "/"
//func HomePage(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "Hello World!",
//	})
//}
//
//// PostHomePage : GET, "/"
//func PostHomePage(c *gin.Context) {
//	body := c.Request.Body
//	value, err := ioutil.ReadAll(body)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	c.JSON(200, gin.H{
//		"message": string(value),
//	})
//}
//
//// QueryStrings : GET, "/query"
//func QueryStrings(c *gin.Context) {
//	name := c.Query("name")
//	age := c.Query("age")
//
//	c.JSON(200, gin.H{
//		"name": name,
//		"age":  age,
//	})
//}
//
//// PathParameters : GET, "/path"
//func PathParameters(c *gin.Context) {
//	name := c.Param("name")
//	age := c.Param("age")
//
//	c.JSON(200, gin.H{
//		"name": name,
//		"age":  age,
//	})
//}

func main() {
	//smtp.EmailSendMain()
	//httpGetPost.PostEx()
	//
	//httpGetPost.JsonPost(struct {
	//	Name string
	//	Age  int
	//}{Name: "Alex", Age: 20})
	//
	//httpGetPost.HttpClient()

	r := gin.Default()
	//r.GET("/", HomePage)
	//r.POST("/", PostHomePage)
	//r.GET("/query", QueryStrings)             // /query?name=brenden&age=20
	//r.GET("/path/:name/:age", PathParameters) // /path/brenden/20

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
