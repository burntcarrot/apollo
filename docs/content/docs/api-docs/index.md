---
title: 'API Documentation'
draft: false
weight: 4
summary: API Documentation for the Apollo server.
---

Apollo Server exposes three endpoints:
- `/api/v1/register`: for registering a service
- `/api/v1/health`: for retrieving health status of all services
- `/api/v1/health/:id`: for retrieving health status of a certain service


## Routes

## `/api/v1/register`

Endpoint for registering a service to the Apollo server.

A service can be registered through:
- Apollo SDKs (for example, `apollo-go-redis`)
- Manually through API clients, etc.

The ideal way is to use Apollo SDKs. Integrate Apollo SDKs into your application to register new services to the Apollo server.

**Request Body Params**:
- `id`: ID for the service
- `name`: Service name
- `uri`: The `/health` endpoint exposed by the service

**Example request**:

```json
{
    "id": "26F5QP6q4UxImHkIr9miSHj3G3U",
    "name": "Service_1",
    "uri": "localhost:9989/health"
}
```

---

## `/api/v1/health`

Endpoint for retrieving health status of services registered to the Apollo server.

Example response:

```json
{
    "results":{
        // TODO:
    }
}
```

---

## `/api/v1/health/:id`

Endpoint for retrieving health status of a **certain** service registered to the Apollo server.

**Params**:
- `id`: The unique ID for the service (KSUID)

Example response:

```json
{
    "results":{
        // TODO:
    }
}
```
