package main

import (
	"context"
        "github.com/gin-gonic/gin"
	"net/http"
)

func getPoints(c *gin.Context) {
  points := []Point{
    Point{5,3},
    Point{3,6},
  }
  c.IndentedJSON(http.StatusOK, points)
}

func getData(c *gin.Context) {
	hiTemplate("rob").Render(context.Background(), c.Writer)
}

func main() {
  listenerMask := "0.0.0.0:3333"
  router := gin.Default()
  router.GET("/points",getPoints)
  router.GET("/data",getData)
  router.Static("/static","./static")
  router.Run(listenerMask)   
}
