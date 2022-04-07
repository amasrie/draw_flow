package main

import (
    "net/http"
    "log"
    "github.com/go-chi/chi/v5"
)

func main() {
    port := "9020"
    log.Printf("Starting up on http://localhost:%s", port)
    route := chi.NewRouter()

    route.Get("/", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "text/plain")
      writer.Write([]byte("Hello World!"))
    })

    log.Fatal(http.ListenAndServe(":" + port, route))
}
