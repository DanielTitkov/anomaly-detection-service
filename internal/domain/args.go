package domain

type (
	FilterAnomaliesArgs struct {
		JobID     int
		Processed *bool
		// TODO: add other fields
	}
	FilterDetectionJobsArgs struct {
		ID     int
		SiteID string
		// TODO: add other fields
	}
)
