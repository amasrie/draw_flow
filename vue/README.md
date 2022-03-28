# Vue Application

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

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

After that, the application could be accessed from `localhost:9200`

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
