package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
type Produce struct {
	Name string `json:"name"`
	Produce_Code string `json:"code"`
	Unit_Price float64 `json:"price"`
}

var prodList = []Produce{
	{
		Produce_Code: "A12T-4GH7-QPL9-3N4M",
		Name: "Lettuce",
		Unit_Price: 3.46,
	},
	{
		Produce_Code: "E5T6-9UI3-TH15-QR88",
		Name: "Peach",
		Unit_Price: 2.99,
	},
	{
		Produce_Code: "YRT6-72AS-K736-L4AR",
		Name: "Green Pepper",
		Unit_Price: 0.79,
	},
	{
		Produce_Code: "TQ4C-VV6T-75ZX-1RMR",
		Name: "Gala Apple",
		Unit_Price: 3.59,
	},
}

func main() {
	r := gin.Default()

	prodRoutes := r.Group("/produce")
	{
		prodRoutes.GET("/", getHandler)
		prodRoutes.POST("/", createItem)
		prodRoutes.DELETE("/:id", deleteItem)
	}	

	if err := r.Run("localhost:4001"); err != nil {
		log.Fatal(err.Error())
	}
}

func getHandler(c *gin.Context) {
	c.JSON(200, prodList)
}

func createItem(c *gin.Context) {
	var reqBody Produce
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error": true,
			"message": "invalid request body",
		})
		return
	}

	reqBody.Produce_Code = uuid.New().String()

	prodList = append(prodList, reqBody)
	c.JSON(200, gin.H{
		"error": false,
	})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, p := range prodList {
		if p.Produce_Code == id {
			prodList = append(prodList[:i], prodList[i + 1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": true,
		"message": "invalid id",
	})
}