# Go Application

If you want to create an isolated Go container, follow these instructions:


```sh
docker run -dit --name go_app -p 9020:9020 -v "$PWD:/go/src" go 
docker start go_app
docker exec -it go_app sh
```

## Project setup

### Create the `go.sum` file, if needed

```
go get nodes
```

### Update dependencies

```
go mod tidy
```

### Run the app for development

```
go run .
```

## Available endpoints

* `/execute`: Post method that requieres a python code. Return the result of executing that code (or error) if applies.
* `/list`: Get method that returns the list of stored graphs.
* `/find`: Get method that requieres a name. Returns the graph associated to the name.
* `/save`: Post method that requieres a JSON with a name and a code. Stores those elements into the database.

## Dependencies

To properly use this service, you will need to create at least two Dgraph nodes (zero and alpha). Optionally, you may also want to add a ratel node if you want a Dgraph UI. After that, make sure the `DGRAPH` variable in the `main.go` file is pointing to the right host of your alpha node (by default, it uses port 9080).


