package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestJsonHolder(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"result": err,
		})
		return
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result["form"])

	c.JSON(200, gin.H{
		"result": result["form"],
	})
	return
}
