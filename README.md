# Table of content
1. [Local development](#local-development)
2. [Dockerized](#dockerized)
3. [Test](#test)

* Change ENV base on .env.example

#Local Development

1. Install go dep [here](https://github.com/golang/dep) and follow instruction for install golang dependency
2. Run main.go file
```
{ENV} go run ./main.go
```
3. Build binary
```
go build
```

# Dockerized

**Docker Compose**

```docker
docker-compose up --build -d
```

# Test

**Run test**

```
{ENV} go test ./tests/...
```
