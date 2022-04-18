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

const DGRAPH = "172.23.0.87:9080"

type code_struct struct {
    Python string `json:"code"`
}

type graph_struct struct {
    Name string `json:"name"`
    Code string `json:"code"`
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
  
    route.Get("/", func(writer http.ResponseWriter, request *http.Request) {
      writer.Header().Set("Content-Type", "text/plain")
      writer.Write([]byte("Hello World!"))
    })

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
        query := `
        {
          getAll(func: has(Code)) {
            Name
          }
        }
        `
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
      var code code_struct
      decoder := json.NewDecoder(request.Body)
      decoder.DisallowUnknownFields()
      err := decoder.Decode(&code)
      if err != nil {
        log.Printf("An error occured while decoding the request body")
        writer.WriteHeader(http.StatusExpectationFailed)
      }
      query := `
      {
        me(func: eq(name, "`+ graphName +`")) {
          Name
          Code
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
      var code graph_struct
      request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
      decoder := json.NewDecoder(request.Body)
      decoder.DisallowUnknownFields()
      err := decoder.Decode(&code)
      if err != nil {
        log.Printf("An error occured while decoding the request body")
        msg = "An error occured while decoding the request body"
        writer.WriteHeader(http.StatusExpectationFailed)
      }
      ctx := context.Background()
      connection := connect()
      if (connection != nil) {
        // transaction := connection.NewTxn()
        // defer transaction.Discard(ctx)
        if err := connection.Alter(ctx, &api.Operation{DropAll: true}); err != nil {
          log.Printf("The drop all operation should have succeeded")
          msg = "The drop all operation should have succeeded"
          writer.WriteHeader(http.StatusExpectationFailed)
        }
        op := &api.Operation{}
        op.Schema = `
          name: string .
          code: string .`
        if err := connection.Alter(ctx, op); err != nil {
          log.Printf("Error on operation")
          msg = "Error on operation"
          writer.WriteHeader(http.StatusExpectationFailed)
        }
        m1 := `
          _:n1 <name> "`+ code.Name +`" .
        `
        mu := &api.Mutation{
          SetNquads: []byte(m1),
          CommitNow: true,
        }
        if _, err := connection.NewTxn().Mutate(ctx, mu); err != nil {
          log.Printf("Error while executing first mutation")
          msg = "Error while executing first mutation"
          writer.WriteHeader(http.StatusExpectationFailed)
        }
        req := &api.Request{CommitNow: true}
        req.Query = `
          query {
            me(func: eq(name, "`+ code.Name +`")) {
              v as uid
            }
          }
        `
        m2 := `uid(v) <code> "`+ code.Code +`" .`
        mu.SetNquads = []byte(m2)
        req.Mutations = []*api.Mutation{mu}
        if _, err := connection.NewTxn().Do(ctx, req); err != nil {
          log.Printf("Error while executing second mutation")
          msg = "Error while executing second mutation"
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
