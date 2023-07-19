package main

import (
	"context"
	"fmt"
	"log"
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

func getData(w http.ResponseWriter, r *http.Request) {
	hiTemplate("rob").Render(context.Background(), w)
}

func main() {
	listenerMask := "0.0.0.0:3333"
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/data", getData)

	log.Printf("Start web server on %s", listenerMask)
	panic(
		fmt.Sprintf(
			"http server crashed: %v",
			http.ListenAndServe(listenerMask, nil),
		),
	)
}
