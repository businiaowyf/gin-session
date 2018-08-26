# gin-session
This is a session middleware for gin web framework. It's based on Gorilla/sessions and use boj/redistore as a redis store. The major thing it does is call context.Clear() to avoid memory leak.

Get Started
---
## Install
```
get get -u github.com/businiaowyf/gin-session/redisess
dep ensure
```

## Build and Run
```
go build
./examples
```

## Test
```
curl http://localhost:8080/test
```

## Note
You need install redis on your local machine first.
