First, create a network to connect all the required containers:

```sh
docker network create ast_nodes
```

Run the following command from this folder to create a Node.js Docker image with Vue CLI:

```sh
docker build -t vue .
```

Crate a Docker container for the Vue.js application and execute the container:

```sh
docker run -dit --name vue_app -p 9200:9021 --network ast_nodes -v "$PWD:/home/vue" vue
docker exec -it vue_app bash
```


docker run -dit --name go_app -p 9201:9021 --network ast_nodes -v "$PWD:/go/src" go
docker exec -it go_app sh

go run .


docker run -it dgraph/dgraph:latest dgraph



