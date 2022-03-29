# Go Application

If you want to create an isolated Go container, follow these instructions:


```sh
docker run -dit --name go_app -p 9020:9020 -v "$PWD:/go/src" go 
docker start go_app
docker exec -it go_app sh
```

## Project setup

### Run the app for development

```
go run .
```

