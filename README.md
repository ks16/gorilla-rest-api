# HTTP API template (`net/http` + `gorilla/mux`)

## How to run
```
docker-compose up
```

## How to test

### Request without `X-Request-ID`
```
curl --location --request GET 'http://localhost:8081/_health'
```

### Request with `X-Request-ID` plus call to Service-B

```
curl --location --request GET 'http://localhost:8081/ping' --header 'X-Request-ID: 139279aa-23bd-4bf8-9a78-d9cb3f34b5a1'
```
