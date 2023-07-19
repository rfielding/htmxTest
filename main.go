package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)


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
