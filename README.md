## MyHTTP

This tool displays the MD5 hash of the response body from the given
list of URLs

### Install

```bash
go get github.com/AgentNemo00/myhttp
```

or clone

```bash
git clone github.com/AgentNemo00/myhttp
```

### Tests

```
go test ./...
```

## Run tests and generate coverage

````shell
go test -v ./... -coverprofile=coverage.out -covermode=count
````

## Generate HTML report

````shell
go tool cover -html=coverage.out -o coverage.html
````

## Generate console report

````shell
go tool cover -func=coverage.out
````

### Usage

```bash
go run cmd/myhttp/main.go example.com
```

or multiple urls by space separated

```bash
go run cmd/myhttp/main.go adjust.com example.com https://httpbin.org
```

you can adjust the parallel processes using the flag

```bash
go run cmd/myhttp/main.go -parallel=2 adjust.com example.com httpbin.org
``` 

the default is 10 process

