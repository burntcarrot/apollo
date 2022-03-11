---
title: 'Redis SDK'
draft: false
weight: 5
summary: Documentation for the Apollo Redis SDK.
---

Apollo Redis SDK is used for performing health checks on Redis and registering services to the Apollo server.

## Quickstart

Create a ID (KSUID) for registering the service to the Apollo server:

```go
id := ksuid.New().String()
```

Set up a service using the generated ID:

```go
c := apollo_go_redis.Setup(id, "http://localhost:9989/health", "http://localhost:8080/api/v1/register")
```

Create a new Redis [checker](#checker) with a strategy:

```go
redisChecker := redis.NewChecker(rc, "set-get")
c.RegisterChecker("redis", redisChecker)
```

Perform checks and get results:

```go
rs := c.PerformChecks()
```

Expose results through the `/health` endpoint:

```go
c.Serve(rs)
```

## Checker

A checker is used for performing health checks on Redis.

Checkers can be configured to use different strategies for health checks. Here is a list of all available strategies:
- **set-get**: Set and get a key.
- **ping**: A simple ping.

## Example

Here is a complete example on how to use the Apollo Redis SDK:

```go
package main

import (
    "fmt"

    apollo_go_redis "github.com/burntcarrot/apollo-go-redis"
    "github.com/burntcarrot/apollo-go-redis/checkers/redis"
    goredis "github.com/go-redis/redis/v8"
    "github.com/segmentio/ksuid"
)

func main() {
    // generate a KSUID for the service
    id := ksuid.New().String()

    // setup (register) a service
    c := apollo_go_redis.Setup(id, "http://localhost:9989/health", "http://localhost:8080/api/v1/register")

    // create Redis connections
    rc := goredis.NewClient(&goredis.Options{
        Addr: "localhost:6379",
    })

    // create a Redis checker using the "set-get" strategy
    redisChecker := redis.NewChecker(rc, "set-get")
    // register the Redis checker
    c.RegisterChecker("redis", redisChecker)

    // perform checks and get results
    rs := c.PerformChecks()

    // expose results to the health endpoint
    c.Serve(rs)
}
```
