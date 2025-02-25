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

### /hello
```bash
open http://localhost:8080/hello
```

You will get:

```json
{
    "latency": "100.583ms",
    "message": "world"
}
```

### /world
```bash
open http://localhost:8080/world
```

You will get:

```json
{
    "latency": "100.123ms",
    "message": "hello"
}
```
