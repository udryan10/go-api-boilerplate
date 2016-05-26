# go-api-boilerplate

### Packages
  - [Gin](https://github.com/gin-gonic/gin) - HTTP Web Framework
  - [Glide](https://github.com/Masterminds/glide) - Vendor package management / version pinning

### Running
#### Clone
`git clone https://github.com/udryan10/go-api-boilerplate.git`
#### Install dependencies
`glide install`
#### Running
```
$ go run main.go

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /status                   --> github.com/udryan10/go-api-boilerplate/controllers.Status (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

#### Testing
```
$ curl http://localhost:8080/status -v

*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /status HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 26 May 2016 02:45:13 GMT
< Content-Length: 3
<
{}
* Connection #0 to host localhost left intact```
