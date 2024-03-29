package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Write a handler to return all items

// assign the handler function to an endpoint path.
func main() {
	router := gin.Default()
	router.GET("/getAlbums", getAlbums)
	router.GET("/getAlbumByID/:id", getAlbumByID)
	router.POST("/postAlbums", postAlbums)
	router.PUT("/putAlbumList", putAlbumList)

	router.LoadHTMLGlob("albumCRUD/templates/*")
	//router.LoadHTMLFiles("web/templates/*")
	//router.SetHTMLTemplate()
	router.GET("/", getIndex)
	router.Run("localhost:8080") // default port : 3000기본
}
func getIndex(c *gin.Context) {
	//// Use a service account
	//ctx := c.Background()
	//sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	//app, err := firebase.NewApp(ctx, nil, sa)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//client, err := app.Database(c)
	//if err != nil {
	//	log.Fatalln("Error initializing database client:", err)
	//}
	//ref := client.NewRef("server/saving-data/fireblog")
	println("getIndex function")
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// serialize the struct into JSON
	// can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON
	c.IndentedJSON(http.StatusOK, albums)
}

// Write a handler to add a new item

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newAlbum)
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//todo : file upload해서 album에 데이터 저장하기
func putAlbumList(c *gin.Context) {
	var newAlbums []album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbums); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newAlbums)
	// Add the new album to the slice.
	for i := range newAlbums {
		albums = append(albums, newAlbums[i])
	}
	c.IndentedJSON(http.StatusCreated, newAlbums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
