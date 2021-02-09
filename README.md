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

## Solution

The service allows to perform one-time and scheduled **detection jobs** for anomaly detection. 

```go
DetectionJob struct {
    ID          int
    Schedule    string // cron string
    Method      string // e.g, 3-sigmas
    SiteID      string
    Metric      string
    Attribute   string
    TimeAgo     string // e.g. 30d
    TimeStep    string // e.g. 1d
    Description string
}
```
Job can be executed right away or scheduled. On job execution service fetches a dataset (mocked by now) and searches for outliers (mocked by now).
Each job execution is saved as **detection job instance** and returns an array of **anomalies**.
**Anomaly** describes a data point marked as outlier. It hodls data points parametres, related job instance id and processed status. The later allow to set anomalies as approved/accepted and stop notifying about them. 

```go
Anomaly struct {
    ID                     int
    DetectionJobID         int     // related job
    DetectionJobInstanceID int     // related instance
    Type                   string  // warning or alarm
    Value                  float64 // outlier item value
    Processed              bool    // if anomaly is accepted/approved
    PeriodStart            time.Time
    PeriodEnd              time.Time
}
```

If anomaly is not *processed* regular notification process will get it and send a notification about it (mocked by now).

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

If **sync** is *false* service only informs that job was created. 

```json
{
    "status": "ok",
    "message": "job created"
}
```

If **sync** is *true* the handler will wait until job is completed and return result.

```json
{
    "Job": {
        "ID": 20,
        "Schedule": "* 10 * 2 1",
        "Method": "3-sigmas",
        "SiteID": "brax",
        "Metric": "foo_metric",
        "Attribute": "bar_attribute",
        "TimeAgo": "30d",
        "TimeStep": "1d",
        "Description": "just a sample job"
    },
    "Instance": {
        "ID": 11,
        "DetectionJobID": 20,
        "StartedAt": "2021-02-09T19:30:35.631405278+03:00",
        "FinishedAt": "2021-02-09T19:30:35.631447879+03:00"
    },
    "Result": [
        {
            "ID": 26,
            "DetectionJobID": 20,
            "DetectionJobInstanceID": 11,
            "Type": "warning",
            "Value": -0.35018186126543377,
            "Processed": false,
            "PeriodStart": "2021-01-25T19:30:35.631428292+03:00",
            "PeriodEnd": "2021-01-25T19:30:35.631428292+03:00"
        }
    ]
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

### /api/v1/listAnomalies

Get list of all anomalies or apply fitler. 

#### Request

```json
{
    "filter": {
        "jobId": 18,
        "processed": false
    }
}
```

**filter** is optional. 
Anomalies can be filtered by **jobId** and **processed** factor.

#### Response 

```json
{
    "Anomalies": [
        {
            "ID": 21,
            "DetectionJobID": 18,
            "DetectionJobInstanceID": 9,
            "Type": "warning",
            "Value": -0.28370364872687986,
            "Processed": false,
            "PeriodStart": "2021-01-12T12:56:54.187577Z",
            "PeriodEnd": "2021-01-12T12:56:54.187577Z"
        }
    ]
}
```
### /api/v1/deleteJob (*not implemented*)

Allows to delete job and remove in from schedule

### /api/v1/setAnomalyStatus (*not implemented*)

Allows to update anomaly status - set processed as *true*/*false*.
NB: if **approveOnSend** option in configs is *true*, **processed** is automatically set to *true* after notifying. 

## Limitations

* use of domain models in API
* no method to inspect jobs history
* distributed jobs are not supported
* job uniqueness is not checked