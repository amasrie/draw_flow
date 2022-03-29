# Draw Flow

## Set Docker environment

Run the following commands to create the Docker images for Vue and Go respectively:

```sh
docker build -t vue -f Dockerfile_vue .
docker build -t go -f Dockerfile_go .
```

After that, just run the doccker-compose file:

```sh
docker-compose -f docker-compose.yml up
```

Then click [here](https://github.com/amasrie/draw_flow/tree/master/go) and [here](https://github.com/amasrie/draw_flow/tree/master/vue) in order to proceed with app setup for Go and Vue apps respectively.
