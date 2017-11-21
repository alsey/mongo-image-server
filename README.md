# MongoDB Image Server

  An image server based mongodb gridfs.

## Installation

```bash
$ go get github.com/alsey/mongo-image-server
```

## Connfiguration

  Use envionment variables to config this server.

- PORT0 : port, default is 3000
- MONGO_ADDR : mongodb address, default is 127.0.0.1:27017
- MONGO_DB : mongodb database name
- MONGO_USER : mongodb authenticate user
- MONGO_PASS : the user's password

## How to Use

  1. While some image files saved at mongodb with gridfs format, you can access this images by URL with this server.

  The URL is

```
http://<host>:<port>/images/<filename>
```

  2. You can change image size by url query parameters while access the image file.

```
http://<host>:<port>/images/<filename>?w=<width>&h=<height>
```

  The url variables `w` and `h` is image's width and height, in pixels.

  3. I wrote a Dockerfile for you, you can make a Docker image, and put it in Kubernetes or Mesos.

```bash
$ docker build -t 'mongo-image-server' .
```

  4. The server includes a health check endpoint for microservice scenario.

```
http://<host>:<port>/health
http://<host>:<port>/env
```

## License

  [MIT](LICENSE)