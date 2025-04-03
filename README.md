# Hello HTTP

A minimal hello-world http server for local test and development.

## Start

```bash
docker run -p 8080:8080 -e LATENCY=100 frwentianqi/hello-http
```

### Configuration 
| Name    | Description                                   | Type   | Sample Value |
|---------|-----------------------------------------------|--------|--------------|
| LATENCY | Add a latency in milliseconds to the response | number | 100          |

> Use environment variable to configure

## Routes

### /status
```bash
open http://localhost:8080/status

{
    "latency": "100.583ms",
    "message": "ok"
}
```

### /version
```bash
open http://localhost:8080/version

{
    "latency": "100.583ms",
    "version": "1.1.0"
    "message": "ok"
}
```

### /hello
```bash
open http://localhost:8080/hello

{
    "latency": "100.583ms",
    "message": "world"
}
```

### /hello/world
```bash
open http://localhost:8080/hello/world

{
    "latency": "100.583ms",
    "message": "hello world"
}
```

### /world
```bash
open http://localhost:8080/world

{
    "latency": "100.123ms",
    "message": "hello"
}
```

### /world/hello
```bash
open http://localhost:8080/world/hello

{
    "latency": "100.123ms",
    "message": "world hello"
}
```
