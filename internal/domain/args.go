package domain

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
)
