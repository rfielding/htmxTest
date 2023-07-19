package main

import (
	"context"
        "github.com/gin-gonic/gin"
	"net/http"
        "log"
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
  log.Printf("note: 127.0.0.1 is the only trusted proxy in this setup")
  listenerMask := "0.0.0.0:3333"
  router := gin.Default()
  router.SetTrustedProxies([]string{"127.0.0.1"})
  router.GET("/points",getPoints)
  router.GET("/data",getData)
  router.Static("/static","./static")
  router.Run(listenerMask)   
}
