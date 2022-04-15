package main

import (
    "net/http"
    "log"
    "encoding/json"
    "os/exec"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

type code_struct struct {
    Python string `json:"code"`
}

func main() {
    port := "9020"
    log.Printf("Starting up on http://localhost:%s", port)
    route := chi.NewRouter()
    route.Use(middleware.Recoverer)
    
    route.Use(cors.Handler(cors.Options{
      AllowedOrigins:   []string{"https://*", "http://*"},
      AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
      AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
      ExposedHeaders:   []string{"Link"},
      AllowCredentials: false,
      MaxAge:           300,
    }))
  
    route.Get("/", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "text/plain")
      writer.Write([]byte("Hello World!"))
    })

    route.Post("/execute", func(writer http.ResponseWriter, request *http.Request) {
      var code code_struct
      request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
      decoder := json.NewDecoder(request.Body)
      decoder.DisallowUnknownFields()
      err := decoder.Decode(&code)
      if err != nil {
        log.Printf("An error occured while decoding the request body")
      }
      cmd := exec.Command("/usr/bin/python3", "-c", code.Python)
      out, err := cmd.CombinedOutput()
        log.Printf(string(out))
      if err != nil {
        log.Printf("An error was found while executing the code")
      }
      writer.Header().Set("Content-Type", "text/plain")
      writer.Write(out)
    })


    http.ListenAndServe(":" + port, route)
}
