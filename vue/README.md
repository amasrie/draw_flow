# Vue Application

If you want to create an isolated Vue container, follow these instructions:

```sh
docker run -dit --name vue_app -p 9021:9021 -v "$PWD:/home/vue" vue
docker start vue_app
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

After that, the application could be accessed from `localhost:9021`

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
