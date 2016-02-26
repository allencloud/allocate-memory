# allocate-memory

Allocate memory is web application supposed to dynamically allocate memory according to web request.

## web API
There is only one API:

```
GET /memory/:size/action/allocate
```

size means memory needed to be allocated with unit of MB

## Listen Port

This application will listen on port 8080.

## Guidance
###1.Start web application

```
cd allocate-memory
export GOPATH=`pwd`
go get -d
go build
./allocate-memory
``` 

###2.test
allocate 256 MB memory
```
curl http://127.0.0.1:8080/memory/256/action/allocate 
```