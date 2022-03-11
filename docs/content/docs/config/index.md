---
title: 'Configuration'
draft: false
weight: 4
summary: Configuring Apollo server through configuration files.
---

Apollo server supports configuration through configuration files.

Configuration files should be in YAML (`.yaml`), for example:

```yaml
redis:
  addr: "localhost:6379"
  password: ""
  db: 0

server:
  addr: "localhost:8080"

logging:
  file: "./test.log"

timeout:
  type: "minute"
  value: 5
```

## Fields

### redis

`redis` is used for Redis configurations.

Available configurations:
- `addr`: Redis address
- `password`: Redis password
- `db`: Redis DB

### server

`server` is used for server-level configurations.

Available configurations:
- `addr`: Server address (example, `localhost:8080` or `http://localhost:8080`)

### logging

`logging` is used for logging configurations.

Available configurations:
- `file`: Relative or absolute path of log file

### timeout

`timeout` is used for timeout related configurations.

Available configurations:
- `type`: Timeout type (`microsecond`, `second`, `minute`)
- `value`: Timeout value
