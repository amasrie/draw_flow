package main

import (
    "net/http"
    "log"
    "encoding/json"
    "os/exec"
    "context"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
    "google.golang.org/grpc"
    "github.com/dgraph-io/dgo/v210"
    "github.com/dgraph-io/dgo/v210/protos/api"
)

const DGRAPH = "192.168.1.107:9080"

type code_struct struct {
    Python string `json:"code"`
}

type Python struct {
    Uid string `json:"uid,omitempty"`
    Name string `json:"name,omitempty"`
    Code string `json:"code,omitempty"`
}

/**
* Creates a connection with a digraph database and closes it on error
*/
func connect() *dgo.Dgraph {
    conn, err := grpc.Dial(DGRAPH, grpc.WithInsecure())
    if err != nil {
      log.Printf("Couldn't establish connection to Dgraph")
      return nil
    } else {
      // defer conn.Close()
      dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
      return dgraphClient
    }
}

/**
* Main function, it's automatically called
*/
func main() {
    port := "9020"
    log.Printf("Starting up on http://localhost:%s", port)
    route := chi.NewRouter()
    
    route.Use(cors.Handler(cors.Options{
      AllowedOrigins:   []string{"https://*", "http://*"},
      AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
      AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
      ExposedHeaders:   []string{"Link"},
      AllowCredentials: false,
      MaxAge:           300,
    }))
  
    /**
    * Gets the list of the already stored graph of code
    */
    route.Get("/list", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "application/json")
      list := []byte("")
      ctx := context.Background()
      connection := connect()
      if (connection != nil) {
        transaction := connection.NewTxn()
        // defer transaction.Discard(ctx)
        query := `{ me(func: has(code)) { name } }`
        resp , err := transaction.Query(ctx, query)
        if err != nil {
          log.Printf("An error occured while trying to get the list of stored graphs")
          writer.WriteHeader(http.StatusExpectationFailed)
        } else {
          list = resp.Json
        }
      }
      writer.Write(list)
    })

    /**
    * Gets a specific code based on the requested id
    */
    route.Get("/find", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "application/json")
      foundElement := []byte("")
      graphName := request.URL.Query().Get("name")
      ctx := context.Background()
      query := `
      {
        me(func: eq(name, "`+ graphName +`")) {
          name
          code
        }
      }
      `
      connection := connect()
      if (connection != nil) {
        transaction := connection.NewTxn()
        // defer transaction.Discard(ctx)
        resp , err := transaction.Query(ctx, query)
        if err != nil {
          log.Printf("An error occured while trying to get a graph item")
          writer.WriteHeader(http.StatusExpectationFailed)
        } else {
          foundElement = resp.Json
        }
      }
      writer.Write(foundElement)
    })

    /**
    * Stores a graph with python code
    */
    route.Post("/save", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "text/plain")
      msg := "Changes saved successfully"
      var code Python
      request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
      decoder := json.NewDecoder(request.Body)
      decoder.DisallowUnknownFields()
      err := decoder.Decode(&code)
      if err != nil {
        msg = "An error occured while decoding the request body"
        log.Printf(msg)
        writer.WriteHeader(http.StatusExpectationFailed)
      }
      ctx := context.Background()
      connection := connect()
      if (connection != nil) {
        // transaction := connection.NewTxn()
        // defer transaction.Discard(ctx)
        if err := connection.Alter(ctx, &api.Operation{DropAll: true}); err != nil {
          msg = "The drop all operation should have succeeded"
          log.Printf(msg)
          writer.WriteHeader(http.StatusExpectationFailed)
        }
        op := &api.Operation{Schema: `
          name: string @index(hash) @upsert .
          code: string .
        `}
        if err := connection.Alter(ctx, op); err != nil {
          msg = "Error on operation"
          log.Printf(msg)
          writer.WriteHeader(http.StatusExpectationFailed)
        }
        req := &api.Request{CommitNow: true}
        req.Query = `
        {
          me(func: eq(name, "` + code.Name + `"))
        }
        `
        pb, err := json.Marshal(Python{Name: code.Name, Code: code.Code})
        if err != nil {
          msg = "Error on marshal operation"
          log.Printf(msg)
          writer.WriteHeader(http.StatusExpectationFailed)
        }
        mu := &api.Mutation{SetJson: pb}
        req.Mutations = []*api.Mutation{mu}
        if _, err := connection.NewTxn().Do(ctx, req); err != nil {
          msg = "Error while executing mutation"
          log.Printf(msg)
          writer.WriteHeader(http.StatusExpectationFailed)
        }
      }
      writer.Write([]byte(msg))
    })

    /**
    * Receives a string with python code and return the results of its execution or an error, if any
    */
    route.Post("/execute", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "text/plain")
      var code code_struct
      request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
      decoder := json.NewDecoder(request.Body)
      decoder.DisallowUnknownFields()
      err := decoder.Decode(&code)
      if err != nil {
        log.Printf("An error occured while decoding the request body")
         writer.WriteHeader(http.StatusExpectationFailed)
     }
      cmd := exec.Command("/usr/bin/python3", "-c", code.Python)
      out, err := cmd.CombinedOutput()
      if err != nil {
        log.Printf("An error was found while executing the code")
        writer.WriteHeader(http.StatusExpectationFailed)
     }
      writer.Write(out)
    })


    http.ListenAndServe(":" + port, route)
}
