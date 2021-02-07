// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnomaliesColumns holds the columns for the "anomalies" table.
	AnomaliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "value", Type: field.TypeFloat64},
		{Name: "processed", Type: field.TypeBool},
		{Name: "detection_job_instance_anomalies", Type: field.TypeInt, Nullable: true},
	}
	// AnomaliesTable holds the schema information for the "anomalies" table.
	AnomaliesTable = &schema.Table{
		Name:       "anomalies",
		Columns:    AnomaliesColumns,
		PrimaryKey: []*schema.Column{AnomaliesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "anomalies_detection_job_instances_anomalies",
				Columns: []*schema.Column{AnomaliesColumns[6]},

				RefColumns: []*schema.Column{DetectionJobInstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DetectionJobsColumns holds the columns for the "detection_jobs" table.
	DetectionJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "schedule", Type: field.TypeString, Nullable: true},
		{Name: "method", Type: field.TypeString},
		{Name: "site_id", Type: field.TypeString},
		{Name: "metric", Type: field.TypeString},
		{Name: "attribute", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// DetectionJobsTable holds the schema information for the "detection_jobs" table.
	DetectionJobsTable = &schema.Table{
		Name:        "detection_jobs",
		Columns:     DetectionJobsColumns,
		PrimaryKey:  []*schema.Column{DetectionJobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DetectionJobInstancesColumns holds the columns for the "detection_job_instances" table.
	DetectionJobInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "value", Type: field.TypeFloat64},
		{Name: "processed", Type: field.TypeBool},
		{Name: "detection_job_instance", Type: field.TypeInt, Nullable: true},
	}
	// DetectionJobInstancesTable holds the schema information for the "detection_job_instances" table.
	DetectionJobInstancesTable = &schema.Table{
		Name:       "detection_job_instances",
		Columns:    DetectionJobInstancesColumns,
		PrimaryKey: []*schema.Column{DetectionJobInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "detection_job_instances_detection_jobs_instance",
				Columns: []*schema.Column{DetectionJobInstancesColumns[6]},

				RefColumns: []*schema.Column{DetectionJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnomaliesTable,
		DetectionJobsTable,
		DetectionJobInstancesTable,
	}
)

func init() {
	AnomaliesTable.ForeignKeys[0].RefTable = DetectionJobInstancesTable
	DetectionJobInstancesTable.ForeignKeys[0].RefTable = DetectionJobsTable
}
