---
title: 'API Documentation'
date: 2019-02-11T19:30:08+10:00
draft: false
weight: 4
summary: API Documentation
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
