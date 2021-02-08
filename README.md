# anomaly-detection-service

## Requirements 

* Golang 
* Postgres or Docker

## Getting started

To run local version with dockerized database:
1. Check /configs/dev.yml
2. Check /deployments/dev/docker-compose.yml
3. `make db` to set up database
4. `make dev` to run service locally with dev config

## API

### /api/v1/addJob

Create one-time or scheduled job.

#### Request

```json
{
    "sync": false,
    "job": {
        "method": "3-sigmas",
        "siteId": "brax",
        "metric": "foo_metric",
        "attribute": "bar_attribute",
        "timeAgo": "30d",
        "timeStep": "1d",
        "schedule": "* 10 * 2 1"
    }
}
```

**schedule** is optional. If provided, job will be run accordingly. Otherwise it will be run once as soon as possible. 

#### Response

If **sync** is *false* service only inform if job was created. 

```json
{
    "status": "ok",
    "message": "job created"
}
```

### /api/v1/listJobs

Get list of all jobs or apply fitler. 

#### Request

```json
{
    "filter": {
        "siteId": "brax",
        "scheduled": true
    }
}
```
**filter** is optional. 
Jobs can be filtered by **id**, **siteId**, and **scheduled** factor (if job has schedule) 

#### Response 

```json
{
    "Jobs": [
        {
            "ID": 8,
            "Schedule": "* 10 * 2 1",
            "Method": "3-sigmas",
            "SiteID": "brax",
            "Metric": "foo_metric",
            "Attribute": "bar_attribute",
            "TimeAgo": "30d",
            "TimeStep": "1d",
            "Description": "just a sample job"
        }
    ]
}
```

## Limitations