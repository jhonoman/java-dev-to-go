# Day 1 - Gin Setup and Basic Routing

## Run

```bash
go mod tidy
go run .
```

Server runs at `http://localhost:8080`.

## Endpoints

- `GET /ping`
- `GET /hello/:name`

## Example

```bash
curl http://localhost:8080/ping
curl http://localhost:8080/hello/Budi
```