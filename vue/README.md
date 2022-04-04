# Vue Application

Application based on [this example](https://github.com/jerosoler/drawflow-vue3-example)

If you want to create an isolated Vue container, follow these instructions:

```sh
docker run -dit --name vue_app -p 9501:9501 -p 9601:9601 -v "$PWD:/home/vue" vue
docker start vue_app
docker exec -it vue_app bash
```

Make sure that the `dev` script inside the `package.json` file is using port `9501` while the file `vite.config.js` is using port `9601`.

## Vue 3 + Vite

This template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

## Recommended IDE Setup

- [VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.volar)

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run dev
```

After that, the application could be accessed from `localhost:9221`

