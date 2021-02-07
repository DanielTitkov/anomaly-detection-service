package domain

import "time"

type (
	FilterAnomaliesArgs struct {
		JobID int
		// TODO: add status etc
	}
	FilterDetectionJobsArgs struct {
		ID     int
		SiteID string
		// TODO: add other fields
	}
	FetchDatasetArgs struct {
		SiteID    string
		Metric    string
		Attribute string
		StartDate time.Time
		EndDate   time.Time
	}
)
