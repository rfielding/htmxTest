package main

import (
  "net/http"
  "fmt"
  "log"
  "io"
  "os"
  "context"
  "github.com/a-h/templ"
)


func getHtml(name string) templ.Component {
  return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
    io.WriteString(w, fmt.Sprintf("<i>hello %s</i>", name))
    return nil
  }
}

func getData(w http.ResponseWriter, r *http.Request) {
  getHtml("rob").Render(context.Background(),os.Stdout)
}

func main() {
  listenerMask := "0.0.0.0:3333"
  fs := http.FileServer(http.Dir("./static"))

  http.Handle("/static/",http.StripPrefix("/static/",fs))
  http.HandleFunc("/data",getData)

  log.Printf("Start web server on %s", listenerMask)
  panic(
    fmt.Sprintf(
      "http server crashed: %v", 
      http.ListenAndServe(listenerMask,nil),
    ),
  )
}
