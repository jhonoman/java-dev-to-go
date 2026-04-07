# 14-Day Gin Learning Curriculum

## Prerequisites
- Comfortable with Go basics: functions, structs, interfaces, error handling
- Understand `go mod`, packages, and how to run a project with `go run .`

## Final Outcomes
- Able to build a simple production-ready REST API with Gin
- Able to handle request validation, middleware, basic JWT authentication, and endpoint testing
- Able to deploy an API to a simple cloud service

## Week 1 - Gin and REST API Foundations

### Day 1 - Gin Setup and Basic Routing
- Topic: Gin installation, project structure, router, HTTP methods (`GET`, `POST`)
- Practice: endpoints `GET /ping`, `GET /hello/:name`
- Outcome: first Gin server is running
- Sample solution: `exercises/gin/day01`

### Day 2 - JSON Requests and Responses
- Topic: `c.JSON()`, `c.Param()`, `c.Query()`, `c.ShouldBindJSON()`
- Practice: simple calculator endpoint (`POST /sum`)
- Outcome: can receive and send JSON correctly

### Day 3 - Route Groups and API Versioning
- Topic: `router.Group("/api/v1")`, handler separation
- Practice: `users` endpoint under `/api/v1/users`
- Outcome: endpoint structure is clean and ready to scale

### Day 4 - Basic Middleware
- Topic: logging middleware, request ID, recovery
- Practice: create custom middleware to log method, path, and latency
- Outcome: understand request flow through the middleware chain

### Day 5 - Input Validation
- Topic: binding tags (`binding:"required"`), email format validation, password length validation
- Practice: register endpoint with request body validation
- Outcome: invalid input is rejected with clear error messages

### Day 6 - Consistent Error Handling
- Topic: standard error response format, status code mapping (`400`, `404`, `500`)
- Practice: `Success` and `Error` response helpers
- Outcome: all endpoints use a consistent response format

### Day 7 - Mini Project 1 (In-Memory CRUD)
- Topic: combine routes, handlers, middleware, and validation
- Practice: `notes` CRUD API without a database
- Outcome: complete a small end-to-end API

## Week 2 - Production Patterns

### Day 8 - Folder Architecture and Layering
- Topic: `handler`, `service`, `repository`, `model`, `middleware`
- Practice: refactor the mini project into a multi-layer structure
- Outcome: codebase is more maintainable

### Day 9 - Database Integration (PostgreSQL)
- Topic: DB connection, basic queries, context timeout
- Practice: move `notes` CRUD from memory to PostgreSQL
- Outcome: API connected to a real database

### Day 10 - Basic JWT Auth
- Topic: login flow, token generation, bearer token auth middleware
- Practice: `POST /login` endpoint and protect `GET /profile`
- Outcome: clear separation between private and public endpoints

### Day 11 - Pagination, Filtering, Sorting
- Topic: standard query params (`page`, `limit`, `q`, `sort`)
- Practice: list endpoint `GET /notes` with pagination
- Outcome: list endpoint is ready for larger datasets

### Day 12 - API Testing
- Topic: `net/http/httptest`, table-driven tests, service mocking
- Practice: test create-note and login endpoints
- Outcome: automated tests for critical endpoints

### Day 13 - Config and Environment
- Topic: `.env`, config struct, dev/prod config separation
- Practice: load port, DB URL, and JWT secret from environment variables
- Outcome: app is ready to run across environments

### Day 14 - Final Project + Deploy
- Topic: basic hardening, graceful shutdown, API documentation
- Practice: final Todo/Auth API project + deploy (Render/Railway/Fly.io)
- Outcome: one portfolio-ready demo project

## Daily Checklist
- Write at least one new endpoint
- Add at least one input validation rule
- Improve at least one part of the code structure
- Record errors you found and how you fixed them

## Final Project Scope (Suggested)
- Auth features: register, login, profile
- Todo features: create, list, detail, update, delete
- Middleware: logging, recovery, JWT auth
- Database: PostgreSQL
- Testing: at least 5 endpoint tests
- Documentation: README + request/response examples

## Next Steps After 14 Days
- Add Redis for caching
- Add rate-limiting middleware
- Use OpenAPI/Swagger
- Learn observability: metrics, structured logging, tracing
