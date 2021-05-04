# Story

I got a fun idea to try out and create a HTTP server which redirects to another
HTTP server which in turn redirects back to the first HTTP server and so on in
a loop or like a deadlock.

I'm using Golang to build this HTTP server

https://duckduckgo.com/?t=ffab&q=golang+http+server&ia=web

https://golangr.com/golang-http-server/

I started using the sample code from the above link

Since I just want to redirect users I'm checking status codes, the 3xx ones and
what header to use, like "Location" to mention the new location

And this has to be a temporary redirect though permanent also will work. I think with permanent - the browser will cache the location. Not sure, gotta check!

https://duckduckgo.com/?t=ffab&q=redirection+http+code&ia=web

https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections

I'm going to get two inputs from the user - listen address which includes port number and then the redirection location, so that I can run two servers like this -

```bash
$ ./redirect-server :8080 localhost:8081
```

```bash
$ ./redirect-server :8081 localhost:8080
```

I'm using `os.Args` to get the program arguments. I'm also using `url` golang built in standard library package to parse the URL to ensure it's a valid one

https://golang.org/pkg/net/url/

https://golang.org/pkg/net/url/#Parse

https://golang.org/pkg/net/url/#URL.String

I got back to the response code stuff. For redirection response code - I'm going to be using `302`

---

I wrote the code and I think I made a big mistake in the user input. The request response cycle looks like this -

```bash
go run main.go :8080 localhost:8081
```

```bash
$ curl -L localhost:8080
curl: (47) Maximum (50) redirects followed

$ curl -v -L localhost:8080
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8080/localhost:8081'
* Found bundle for host localhost: 0x7f8ca2518130 [can pipeline]
* Could pipeline, but not asked to!
* Re-using existing connection! (#0) with host localhost
* Connected to localhost (::1) port 8080 (#0)
> GET /localhost:8081 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: localhost:8081
< Date: Tue, 04 May 2021 16:07:06 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Maximum (50) redirects followed
curl: (47) Maximum (50) redirects followed
* Closing connection 0
```

I think I need to include the `http` scheme in the redirection location

---

I changed the user input

```bash
$ go run main.go :8080 http://localhost:8081
```

and now it works fine

```bash
$ curl -v -L localhost:8080
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Location: http://localhost:8081
< Date: Tue, 04 May 2021 16:09:43 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/'
* Found bundle for host localhost: 0x7ffd85417a30 [can pipeline]
* Could pipeline, but not asked to!
*   Trying ::1...
* TCP_NODELAY set
* Connection failed
* connect to ::1 port 8081 failed: Connection refused
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connection failed
* connect to 127.0.0.1 port 8081 failed: Connection refused
* Failed to connect to localhost port 8081: Connection refused
* Closing connection 1
curl: (7) Failed to connect to localhost port 8081: Connection refused
* Closing connection 0
```

---

I'm going to run one more server and then try it out

```bash
$ go run main.go :8080 http://localhost:8081
```

```bash
$ go run main.go :8081 http://localhost:8080
```

```bash
$ curl -L localhost:8081
curl: (47) Maximum (50) redirects followed
```

Same looping issue. I'm going to try in my browser too!

Firefox tells me

```
The page isn’t redirecting properly

Firefox has detected that the server is redirecting the request for this address in a way that will never complete.

    This problem can sometimes be caused by disabling or refusing to accept cookies.

```

Interesting problem ;)

Anyways, I guess I'm done :P
