package domain

type (
	Anomaly struct {
		ID                     int
		DetectionJobID         int     // related job
		DetectionJobInstanceID int     // related instance
		Type                   string  // warning or alarm
		Value                  float64 // outlier item value
		Processed              bool    // if anomaly is accepted/approved
	}
	DetectionJob struct {
		ID        int
		Schedule  string // cron string
		Method    string // e.g, 3-sigmas
		SiteID    string
		Metric    string
		Attribute string
	}
	DetectionJobInstance struct {
		ID             int
		DetectionJobID int // related job
	}
)
