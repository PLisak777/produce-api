package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// produce represents data about a piece of produce
type produce struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Produce_Code string `json:"code"`
	Unit_Price float64 `json:"price"`
}

// produce sliced to seed list data
var prodList = []produce{
	{
		ID: 1,
		Produce_Code: "A12T-4GH7-QPL9-3N4M",
		Name: "Lettuce",
		Unit_Price: 3.46,
	},
	{
		ID: 2,
		Produce_Code: "E5T6-9UI3-TH15-QR88",
		Name: "Peach",
		Unit_Price: 2.99,
	},
	{
		ID: 3,
		Produce_Code: "YRT6-72AS-K736-L4AR",
		Name: "Green Pepper",
		Unit_Price: 0.79,
	},
	{
		ID: 4,
		Produce_Code: "TQ4C-VV6T-75ZX-1RMR",
		Name: "Gala Apple",
		Unit_Price: 3.59,
	},
}

func main() {
	router := gin.Default()
	router.GET("/produce", getProduce)
	router.POST("/produce", postItems)

	router.Run("localhost:8080")
}

// getProduce responds with list of all produce in JSON
func getProduce(c *gin.Context) {
	c.JSON(http.StatusOK, prodList)
}

//postItems adds produce from JSON received in req body
func postItems(c *gin.Context) {
	var newItem produce

	// Call BindJSON to bind the received JSON to newItem
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	// Add new item to the slice
	prodList = append(prodList, newItem)
	c.JSON(http.StatusCreated, newItem)
}