---
title: 'SDK Documentation'
date: 2019-02-11T19:30:08+10:00
draft: false
weight: 4
summary: SDK Documentation
---

Apollo Server exposes three endpoints:
- `/api/register`: for registering a service
- `/api/health`: for retrieving health status of all services
- `/api/health/:id`: for retrieving health status of a certain service


## Routes

## `/api/register`

**Request Params**:
- `id`: ID for the service
- `name`: Service name
- `uri`: The `/health` endpoint exposed by the service

**Example request**:

```json
{
    "id": 1,
    "name": "Service_1",
    "uri": "localhost:9989/health"
}
```

## `/api/health`

Example response:

```json
{
    "results":{
        // TODO:
    }
}
```

## Usage

```go
package main

import (
	apollo_go_redis "github.com/burntcarrot/apollo-go-redis"
	"github.com/burntcarrot/apollo-go-redis/checkers/redis"
)

func main() {
	c := apollo_go_redis.Setup(1, "http://localhost:9989/health", "http://localhost:8080/api/register")
	redisChecker := redis.NewChecker("localhost:6375", "set-get")
	c.RegisterChecker("redis", redisChecker)

	redisCheckerNew := redis.NewChecker("localhost:6379", "ping")
	c.RegisterChecker("redis-new", redisCheckerNew)

	rs := c.PerformChecks()
	c.Serve(rs)
}

```
