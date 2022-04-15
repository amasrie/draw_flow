# Draw Flow

## Set Docker environment

Run the following command in order to create the Docker images for Dgraph, Vue and Go respectively:


```sh
docker-compose -f docker-compose.yml up
```

Note: make sure that a new volume has been created, otherwise Dgraph may throw a manifest version error.

Click [here](https://github.com/amasrie/draw_flow/tree/master/go) and [here](https://github.com/amasrie/draw_flow/tree/master/vue) to read more about the Go and Vue applications respectively.
